## Env Session

Env Session is a simple Go app that allows you to create a new powershell session with custom command aliases.
I've created this app to make it easier for me to run my apps in a remote environment using a USB drive.

### Building the app
```bash
$ go build
```

### Running

#### First run

Running the app will create a `envs.json` file with a found powershell installation (Only Powershell 7 has been tested).

#### Adding a new command
It's very simple to add a new command to the app. Just add a new entry to the `envs.json` file.

```json
{
    "command_name": "path/to/executable"
}
```

Running the app will open powershell in the default environment (On Windows 11 it is Windows Terminal by default) with the new command aliases available.

### Side Note
There are a few bugs which I'm very clueless on how to fix, this is my first time working with Go :p
