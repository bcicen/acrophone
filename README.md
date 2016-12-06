# acrophone

Commandline utility to easily convert text to phonetic spelling

## Installing

Download the latest release for your platform:

#### OS X

```bash
curl -Lo acrophone https://github.com/bcicen/acrophone/releases/download/v0.3/acrophone-0.3-darwin-amd64
sudo mv acrophone /usr/local/bin/
sudo chmod +x /usr/local/bin/acrophone
```

#### Linux

```bash
wget https://github.com/bcicen/acrophone/releases/download/v0.3/acrophone-0.3-linux-amd64 -O acrophone
sudo mv acrophone /usr/local/bin/
sudo chmod +x /usr/local/bin/acrophone
```

## Usage

Simply provide one or more words as an argument:
```bash
$ acrophone hello
Hotel Echo Lima Lima Oscar
```

## Options

Option | Description
--- | ---
--alphabet <value>, -a <value> |  Phonetic alphabet to use (default: "nato")
--list, -l | List available phonetic alphabets
--version, -v | Print the version
