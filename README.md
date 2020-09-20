# gcectrl

## About

A command line tool to manage Google Compute Engine.

## Preapre

### Setup GCP

* Create Service Acccount
* Create JSON Key and download it.

### Install gcectrl

```sh
wget https://github.com/oreno-tools/gcectrl/releases/download/latest/gcectrl_darwin_amd64 -O ~/bin/gcectrl
chmod +x ~/bin/gcectrl
```

## Usage

### Get instances list

```sh
env GOOGLE_APPLICATION_CREDENTIALS=credential.json gcectrl \
  -project=xxxxxxxxx -zone=xxxxxxx
```

The ouput looks like this:

```sh
+---------------------+----------------+-------------+------------+------------+
|         ID          |      NAME      | MACHINETYPE | IP ADDRESS |   STATUS   |
+---------------------+----------------+-------------+------------+------------+
| xxxxxxxxxxxxxxxxxxx | test-instance1 | f1-micro    | 10.xxx.x.2 | TERMINATED |
+---------------------+----------------+-------------+------------+------------+
```

### Start instance

```sh
env GOOGLE_APPLICATION_CREDENTIALS=credential.json gcectrl \
  -project=xxxxxxxxx -zone=xxxxxxx -start -instance=test-instance1
```

### Stop instance

```sh
env GOOGLE_APPLICATION_CREDENTIALS=credential.json gcectrl \
  -project=xxxxxxxxx -zone=xxxxxxx -stop -instance=test-instance1
```

## Tips

### Please use direnv

Create a `.envrc` file as follows.

```sh
export GOOGLE_APPLICATION_CREDENTIALS=credential.json
export GCP_PROJECT=xxxxxxxxx
```

And run gcectrl.

```sh
gcectrl
```

It's very simply.

### Automatic execution

There is no batch mode in gcectrl, but it is possible to run it automatically.

```sh
echo 'y' | /path/to/gcectrl -start -instance=test-instance1
```
