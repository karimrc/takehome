# Take home assignment

Lets say you have a webserver that acts as a generic command runner. This task runner can take in any linux command,
followed by some args, and return the output of that command after it is done running it. This works great for tasks that finish
within the lifespan of a request, but what about tasks that take longer than that? We need to come up with a solution that allows
users to launch tasks, using our task runner, that can be running for minutes, hours, days, or possibly infinitely. Ideally, we want
to see the progress of the task as it's running by having some sort of access to the output it produces through stdout and stderr.
We would also like to have the ability to end long running tasks ourselves. Keeping a history of all run tasks would be a nice bonus too.
Any other suggestions to making this better would be awesome too. Feel free to use any language you are comfortable with, but we have
included a starter in go in this repo. Making several new api endpoints that take in params to start or monitor a long lived task is a
good starting point.

## API Usage

**command**:

_curl -X POST -H 'Content-Type: application/json' -d '{"command": "python3", "args": ["-c", "print(\"hello world\")"]}' localhost:8080/run_

**output**:

hello world

**command**:

_curl -X POST -H 'Content-Type: application/json' -d '{"command": "cat", "args": ["README.md"]}' localhost:8080/run_

**output**:

`\# Take home assignment

Lets say you have a webserver that acts as a generic command runner. This task runner can take in any linux command,
...`

## Requirements:

    1. Support Long-Running Tasks: Allow the task runner to handle tasks that may run for minutes, hours, days, or infinitely.
    2. Real-Time Output Access: Enable access to the progress of the task while it is running by exposing stdout and stderr outputs.
    3. Task Termination: Provide the ability to terminate long-running tasks.
    4. Task History: Keep a history of all run tasks.
    5. API Endpoints: Create API endpoints to:

       - Launch tasks.
       - Monitor the progress of running tasks.
       - Terminate tasks via ID.
       - Secure API with a token.
