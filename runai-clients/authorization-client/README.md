# Authorization-client

In charge of calling the authorization service evaluation requests based on user context and requested resource permissions.

In order to use this package in your golang control-plane microservice, install by the following steps:
1. in you go.mod add:

        replace github.com/dsx-ai-factory/run-ai/runai-clients/authorization-client => ../../runai-clients/authorization-client

        replace github.com/dsx-ai-factory/run-ai/runai-clients/authorization-client => ../../runai-clients/authorization-client
2. go get github.com/dsx-ai-factory/run-ai/runai-clients/authorization-client
3. go mod tidy
4. To your config file add an Authorization section taken from authorization_client package, and add the value of the authorization url, probably "runai-backend-authorization.runai-backend.svc:8080"
5. Add to the service init function a line to init the client 

   
      authorization_client.Init(conf.Authorization)
6. For each exposed API, you should define its requested resource or mark it as not authorized end point. This happens under the method with the following parameters: (Will work only if you are using the RUNAI shared mustache files for swagger code generation)
 #### For not authorized endpoint (for example health checks), add the parameter


    get:
      x-no-authorization: true


#### For authorized endpoint that does not require specific resourceId permissions, add the parameter


      get:
        x-authorize:
          resource-type: <ResourceType>
          action: <Action>


#### For authorized endpoint that does require specific resourceId permissions, add the parameter


      get:
        x-authorize:
          resource-type: <ResourceType>
          action: <Action>
          resource-id: <IdParamName>

