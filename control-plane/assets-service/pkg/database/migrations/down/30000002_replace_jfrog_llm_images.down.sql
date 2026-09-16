-- SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
-- SPDX-License-Identifier: Apache-2.0

-- Rollback of 20250006_replace_jfrog_llm_images:
-- 1. Revert the llm-server environment image back to the jfrog vLLM build.
-- 2. Remove the per-tenant Legacy gpt2 copies (environment + cloned GPT-2 model)
--    created by the up migration.
-- 3. Restore the gpt2 environment in the fictive tenant and the system-scoped
--    GPT-2 model, matching the state prior to the up migration.

DO $$
DECLARE
    gpt2_env_id       uuid := 'dbf64cad-5652-4698-b24c-b7193230acbe';
    gpt2_model_id     uuid := '608af4f4-1de5-490a-ab14-3fff85111e52';
    fictive_tenant_id int  := 2147483647;
    CREATED_BY        text := 'Run:ai';
    CREATED_AT        timestamp := now();
    IMAGE_PULL_POLICY text := 'IfNotPresent';

    gpt2_env_vars jsonb := '[{"name": "PATH_PREFIX", "value": "/${RUNAI_PROJECT}/${RUNAI_JOB_NAME}", "description": "If the cluster uses path-based URLs (e.g., http://my.domain.com/proj/workload), leave the value as is.\n\nIf the cluster uses host-based URLs (e.g., http://proj-workload.my.domain.com), remove this environment variable entirely."}]';
    gpt2_connections jsonb := '[{"name": "Chatbot UI", "internalToolInfo": {"toolType": "chatbot-ui", "containerPort": 3000, "connectionType": "ExternalUrl", "externalUrlInfo": {}}}, {"name": "ServingPort", "internalToolInfo": {"toolType": "custom", "containerPort": 3000, "connectionType": "ServingPort", "servingPortInfo": {"protocol": "http"}}}]';
    gpt2_model_description text := 'OpenAI''s GPT2 Large is the 774M parameter version of GPT-2, a transformer-based language model created and released by OpenAI. The model is a pre-trained model on English language using a causal language modeling (CLM) objective.';
    rec RECORD;
BEGIN
    -- 1. Revert the llm-server image back to jfrog.
    UPDATE assets_environment ae
    SET image = 'runai.jfrog.io/core-llm/runai-vllm:v0.6.4-0.10.0'
    FROM assets_metadata am
    WHERE ae.asset_id = am.id
      AND am.kind = 'environment'
      AND am.name = 'llm-server'
      AND ae.image = 'vllm/vllm-openai:v0.6.4';

    -- 2. Delete the per-tenant Legacy gpt2 copies (and their cloned GPT-2 models).
    FOR rec IN
        SELECT id FROM assets_metadata
        WHERE name = 'gpt2' AND kind = 'environment'
          AND tenant_id IS NOT NULL AND tenant_id != fictive_tenant_id
    LOOP
        DELETE FROM assets_metadata
        WHERE kind = 'model'
          AND id IN (SELECT asset_id FROM assets_model WHERE environment_asset_id = rec.id);
        DELETE FROM assets_model WHERE environment_asset_id = rec.id;
        DELETE FROM assets_supported_workload_types WHERE asset_id = rec.id;
        DELETE FROM assets_environment WHERE asset_id = rec.id;
        DELETE FROM assets_metadata WHERE id = rec.id;
    END LOOP;

    -- 3. Restore the gpt2 environment (fictive tenant) and the system GPT-2 model.
    BEGIN
        INSERT INTO assets_metadata ("id", "kind", "scope", "tenant_id", "name", "created_by", "created_at", "updated_by", "updated_at")
        VALUES (gpt2_env_id, 'environment', 'tenant', fictive_tenant_id, 'gpt2', CREATED_BY, CREATED_AT, CREATED_BY, CREATED_AT);

        INSERT INTO assets_environment ("asset_id", "image", "image_pull_policy", "environment_variables", "connections")
        VALUES (gpt2_env_id, 'runai.jfrog.io/core-llm/quickstart-inference:gpt2-cpu', IMAGE_PULL_POLICY, gpt2_env_vars, gpt2_connections);

        INSERT INTO assets_supported_workload_types ("asset_id", "workspace", "training", "distributed", "inference")
        VALUES (gpt2_env_id, FALSE, FALSE, FALSE, TRUE);

        INSERT INTO assets_metadata ("id", "kind", "scope", "name", "created_by", "created_at", "updated_by", "updated_at")
        VALUES (gpt2_model_id, 'model', 'system', 'GPT-2', CREATED_BY, CREATED_AT, CREATED_BY, CREATED_AT);

        INSERT INTO assets_model ("asset_id", "title", "description", "thumbnail", "huggingface_id", "license", "environment_asset_id")
        VALUES (gpt2_model_id, 'GPT-2', gpt2_model_description, '', 'openai-community/gpt2', 'runai internal', gpt2_env_id);
    EXCEPTION
        WHEN unique_violation THEN
            RAISE NOTICE 'gpt2 preset already exists, skipping restore';
    END;
END $$;
