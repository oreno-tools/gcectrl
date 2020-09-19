# gcectrl

## About

A command line tool to manage Google Compute Engine.

## Preapre

* Create Service Acccount
* Create JSON Key and download it.

## Usage

### Get Instances

```sh
env GOOGLE_APPLICATION_CREDENTIALS=credential.json gcectrl -project=xxxxxxxxx -zone=xxxxxxx
```

### Start Instance

```sh
env GOOGLE_APPLICATION_CREDENTIALS=credential.json gcectrl -project=xxxxxxxxx -zone=xxxxxxx -start -instance=xxxxxxx
```

### Stop Instance

```sh
env GOOGLE_APPLICATION_CREDENTIALS=credential.json gcectrl -project=xxxxxxxxx -zone=xxxxxxx -stop -instance=xxxxxxx
```

## Tips

Create a `.envrc` file as follows.

```sh
export GOOGLE_APPLICATION_CREDENTIALS=credential.json
export GCP_PROJECT=xxxxxxxxx
```

And run gcectrl.

```sh
gcectrl
```
