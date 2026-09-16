-- SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
-- SPDX-License-Identifier: Apache-2.0

-- Migration to remove references to the private jfrog registry from the built-in
-- environment presets.
--
-- 1. "llm-server": repoint from runai.jfrog.io/core-llm/runai-vllm to the public
--    upstream vLLM OpenAI-compatible server image (covers the fictive-tenant
--    original and any per-tenant provisioned copies).
--
-- 2. "gpt2": deprecate the preset following the same strategy as
--    20250004_deprecate_chatbot_ui:
--      a. For each real tenant that has preset assets, clone gpt2 as a per-tenant copy
--      b. Mark those cloned copies as (Legacy)
--      c. Delete the original from the fictive tenant so new tenants won't get it
--    Unlike chatbot-ui, gpt2 has a GPT-2 model asset bound to it via a FK
--    (ON DELETE RESTRICT), so the model is cloned per-tenant alongside the
--    environment and the fictive/system original is removed last.

DO $$
DECLARE
    gpt2_env_id       uuid := 'dbf64cad-5652-4698-b24c-b7193230acbe';
    gpt2_model_id     uuid := '608af4f4-1de5-490a-ab14-3fff85111e52';
    fictive_tenant_id int  := 2147483647;
    legacy_env_description   text := '(Legacy) GPT-2 quickstart inference environment. This preset is deprecated; consider using the llm-server environment instead.';
    legacy_model_description text := '(Legacy) OpenAI GPT-2 demo model. This preset is deprecated.';
    rec RECORD;
    new_env_id uuid;
    new_model_id uuid;
    env_found boolean := false;
    wt_found boolean := false;
    model_found boolean := false;
    v_image text;
    v_image_pull_policy text;
    v_environment_variables jsonb;
    v_connections jsonb;
    v_workspace boolean;
    v_training boolean;
    v_distributed boolean;
    v_inference boolean;
    m_title text;
    m_description text;
    m_thumbnail bytea;
    m_huggingface_id text;
    m_license bytea;
    m_compute_asset_id uuid;
BEGIN
    -- 1. Repoint the llm-server environment off jfrog.
    UPDATE assets_environment ae
    SET image = 'vllm/vllm-openai:v0.6.4'
    FROM assets_metadata am
    WHERE ae.asset_id = am.id
      AND am.kind = 'environment'
      AND am.name = 'llm-server'
      AND ae.image = 'runai.jfrog.io/core-llm/runai-vllm:v0.6.4-0.10.0';

    -- 2. Deprecate the gpt2 preset (only if it still exists in the fictive tenant).
    IF EXISTS (SELECT 1 FROM assets_metadata WHERE id = gpt2_env_id AND tenant_id = fictive_tenant_id) THEN

        -- Environment spec to clone.
        SELECT image, image_pull_policy, environment_variables, connections
        INTO v_image, v_image_pull_policy, v_environment_variables, v_connections
        FROM assets_environment WHERE asset_id = gpt2_env_id;
        env_found := FOUND;

        -- Supported workload types to clone.
        SELECT workspace, training, distributed, inference
        INTO v_workspace, v_training, v_distributed, v_inference
        FROM assets_supported_workload_types WHERE asset_id = gpt2_env_id;
        wt_found := FOUND;

        -- GPT-2 model spec to clone (bound to the fictive gpt2 environment).
        SELECT title, description, thumbnail, huggingface_id, license, compute_asset_id
        INTO m_title, m_description, m_thumbnail, m_huggingface_id, m_license, m_compute_asset_id
        FROM assets_model WHERE asset_id = gpt2_model_id;
        model_found := FOUND;

        -- For each real, provisioned tenant that doesn't already have its own gpt2,
        -- create per-tenant Legacy copies of the environment (and its GPT-2 model).
        FOR rec IN
            SELECT DISTINCT tenant_id
            FROM assets_metadata
            WHERE tenant_id IS NOT NULL
              AND tenant_id != fictive_tenant_id
              AND kind = 'environment'
              AND tenant_id NOT IN (
                  SELECT tenant_id FROM assets_metadata
                  WHERE name = 'gpt2' AND kind = 'environment'
                    AND tenant_id IS NOT NULL AND tenant_id != fictive_tenant_id
              )
        LOOP
            new_env_id := gen_random_uuid();

            -- Clone environment metadata with Legacy description.
            INSERT INTO assets_metadata (id, kind, scope, tenant_id, name, created_by, created_at, updated_by, updated_at, description)
            SELECT new_env_id, kind, scope, rec.tenant_id, name, created_by, now(), created_by, now(), legacy_env_description
            FROM assets_metadata
            WHERE id = gpt2_env_id
            ON CONFLICT DO NOTHING;

            -- Clone environment spec.
            IF env_found THEN
                INSERT INTO assets_environment (asset_id, image, image_pull_policy, environment_variables, connections)
                VALUES (new_env_id, v_image, v_image_pull_policy, v_environment_variables, v_connections)
                ON CONFLICT DO NOTHING;
            END IF;

            -- Clone workload supported types.
            IF wt_found THEN
                INSERT INTO assets_supported_workload_types (asset_id, workspace, training, distributed, inference)
                VALUES (new_env_id, v_workspace, v_training, v_distributed, v_inference)
                ON CONFLICT DO NOTHING;
            END IF;

            -- Clone the GPT-2 model per-tenant, pointing at the cloned environment.
            IF model_found THEN
                new_model_id := gen_random_uuid();

                INSERT INTO assets_metadata (id, kind, scope, tenant_id, name, created_by, created_at, updated_by, updated_at, description)
                SELECT new_model_id, kind, 'tenant', rec.tenant_id, name, created_by, now(), created_by, now(), legacy_model_description
                FROM assets_metadata
                WHERE id = gpt2_model_id
                ON CONFLICT DO NOTHING;

                INSERT INTO assets_model (asset_id, title, description, thumbnail, huggingface_id, license, compute_asset_id, environment_asset_id)
                VALUES (new_model_id, m_title, m_description, m_thumbnail, m_huggingface_id, m_license, m_compute_asset_id, new_env_id)
                ON CONFLICT DO NOTHING;
            END IF;
        END LOOP;

        -- Delete the fictive/system originals: the GPT-2 model first (FK
        -- ON DELETE RESTRICT), then the gpt2 environment.
        DELETE FROM assets_model WHERE asset_id = gpt2_model_id;
        DELETE FROM assets_metadata WHERE id = gpt2_model_id;
        DELETE FROM assets_supported_workload_types WHERE asset_id = gpt2_env_id;
        DELETE FROM assets_environment WHERE asset_id = gpt2_env_id;
        DELETE FROM assets_metadata WHERE id = gpt2_env_id;
    END IF;

    -- Also mark any pre-existing per-tenant gpt2 copies as Legacy.
    UPDATE assets_metadata
    SET description = legacy_env_description
    WHERE name = 'gpt2' AND kind = 'environment'
      AND tenant_id IS NOT NULL AND tenant_id != fictive_tenant_id;
END $$;
