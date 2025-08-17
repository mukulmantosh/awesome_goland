# Demo Walkthrough


## Junie with MCP


### Prerequisites

- Make sure Docker & npx are already running.
- Create a MongoDB container  

Run the following command to start MongoDB:

```shell
docker run -d \
--name mongodb-container \
-e MONGO_INITDB_ROOT_USERNAME=admin \
-e MONGO_INITDB_ROOT_PASSWORD=password \
-e MONGO_INITDB_DATABASE=mydb \
-p 27017:27017 \
mongo
```

Follow the following steps to establish MCP connection:

![demo4](./images/demo4.png)

![demo5](./images/demo5.png)

![demo6](./images/demo6.png)

Add the following JSON, which includes both Playwright and MongoDB:

* Note: Make sure [npx](https://docs.npmjs.com/cli/v8/commands/npx) is installed in your machine. 

```shell
{
  "mcpServers": {
    "MongoDB": {
      "command": "docker",
      "args": [
        "run",
        "--rm",
        "-i",
        "-e",
        "MDB_MCP_CONNECTION_STRING=mongodb://admin:password@host.docker.internal:27017/mydb?authSource=admin",
        "mongodb/mongodb-mcp-server:latest"
      ]
    },
    "playwright": {
      "command": "npx",
      "args": [
        "@playwright/mcp@latest"
      ]
    }
  }
}
```
Make sure all the connections are already pre-configured before starting the demo. 

### During the Demo

- Create an empty project
- Open Junie, which is placed on the right and run the following prompts.
- Please check `prompts.txt`