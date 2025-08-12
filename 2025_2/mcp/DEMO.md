# Demo Walkthrough


### Junie with MCP



Follow the below steps:

- Create an empty project
- Open Junie, which is placed on the right and run the following prompts.
- Please check `prompts.txt`

**Notes**
- Make sure Docker & npx are already running.
- Create a MongoDB container  

Run the following command to start MongoDB:

```shell

```



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