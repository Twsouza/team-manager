# Team Manager

RESTful API that will help you manage your team. You can create a member (employee or contractor) and attach a tag to him.

## Development

If you use VSCode you get started easily using the extension [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers).

Or you can start the project manually running the `docker-compose.yml`:

```
docker-compose -f .devcontainer/docker-compose.yml up -d
```

You can use the application on [http://127.0.0.1:3000](http://127.0.0.1:3000).

### Swagger

When changes are made, you have to init the docs by running the following command on the terminal:

```
swag init --parseDependency --parseInternal -g actions/app.go
```

## How to use the application

Run the application, then open the documentation on http://localhost:3000/v1/doc/index.html. All endpoint are available for test.

## How to deploy

Follow the steps to [install the heroku CLI](https://devcenter.heroku.com/articles/heroku-cli).

Set the stack to container:

```
heroku stack:set container
```

To deploy the application, push the modification to the Heroku git repository.
