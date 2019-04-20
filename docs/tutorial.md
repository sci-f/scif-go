# Docker Tutorial

Here is a quick tutorial for using the scif-go library, provided in a container
with a filesystem ready to go! If you want to see the equivalent tutorial with the
Python client base, see [this quickstart](https://sci-f.github.io/tutorial-really-quick-start).

## 1. Get container with a scientific filesystem

```bash
$ docker pull vanessa/scif-go:hello-world
```

## 2. View the scientific filesystem entrypoint

```bash
$ docker run vanessa/scif-go:hello-world
```

## 3. Discover Installed Apps

```bash
$ docker run vanessa/scif-go:hello-world apps
```

## 4. Commands

### Interactive shell

```bash
$ docker run -it vanessa/scif-go:hello-world shell
```

### Shell with application active

```bash
$ docker run -it vanessa/scif-go:hello-world shell hello-world-env
```

### Execute

```bash
$ docker run vanessa/scif-go:hello-world exec hello-world-echo echo "Another hello!"
```

### Execute command with environment variable $OMG
```
$ docker run vanessa/scif-go:hello-world exec hello-world-env echo [e]OMG
```

### Run

```bash
$ docker run vanessa/scif-go:hello-world run hello-world-echo
```

### Help

```bash
$ docker run vanessa/scif-go:hello-world help hello-world-env
```

### Inspect

```bash
$ docker run vanessa/scif-go:hello-world inspect --environment hello-world-env
```

### Test

```bash
# Passing Test (test script returns 0 with no arguments)
$ docker run vanessa/scif-go:hello-world test hello-world-script
echo $?

# Failing Test (test script returns argument as return code)
docker run vanessa/scif-go:hello-world test hello-world-script 255
echo $?
```
