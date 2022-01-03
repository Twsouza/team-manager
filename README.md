# Team Manager

RESTful API that will help you manage your team. You can create a member (employee or contractor) and attach a tag to him.

## Development

If you use VSCode you get started easily using the extension [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers).

Or you can start the project manually running the `docker-compose.yml`:

```
docker-compose -f .devcontainer/docker-compose.yml up -d
```

You can use the application on [http://127.0.0.1:3000](http://127.0.0.1:3000).

## Instructions

The goal of this exercise is to create a demo RESTful application using Golang.

### The Task

In this task, we are building the backend of an application that helps us managing our team.

#### Features and Requirements
- A member has a name and a type the late one can be an employee or a contractor - if it's a contractor, the duration of the contract needs to be saved, and if it's an employee we need to store their role, for instance: Software Engineer, Project Manager and so on.
- A member can be tagged, for instance: C#, Angular, General Frontend, Seasoned Leader and so on. (Tags will likely be used as filters later, so keep that in mind)

We need to offer a RESTful CRUD for all the information above.

#### Notes:

1. You can use any Golang framework
2. Make sure to provide a tutorial on how to run your application
3. Feel free to use any database

#### Bonus points:

1. If you add automated tests
2. If you provide a Docker image

## F.A.Q.

### How do you evaluate the exercise?
Our evaluation is based on many aspects, such as general approach adopted, quality of code, use of best practices, capabilities to keep the code simple and maintainable.

### How can I deliver the exercise?
To deliver the exercise, you should clone this repository and work on a new branch. When you'll consider it completed, just push the branch and open a Pull Request.
