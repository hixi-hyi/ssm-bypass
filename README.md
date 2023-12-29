# ssm-bypass
This binary will export the path of SSM ParameterStore as an environment variables.
For example, it can be used in entrypoint.sh.

The following are the registered contents of ssm-parameter and the values output by ssm-bypass.
* ssm-parameter
```
/your/path/HOGE → hoge
/your/path/FUGA → fuga
```

* `ssm-bypass /your/pass`
```
export HOGE=hoge
export FUGA=fuga
```

# usage
1. copy binary for your system from `bin/`
2. write the bellow code into your Dockerfile
```
ADD ssm-bypass_{your_operation_system} ssm-bypass
RUN chmod +x ssm-bypass
```
3. write the bellow code into your entrypoint.sh
```
$(ssm-bypass /your/path/)
```


# caution

When working with environment variables and SSM parameters in this system, please be aware of the following:

- **Special Characters in SSM Parameters**: If the value of an SSM parameter contains shell special characters such as single quotes, double quotes, back quotes, multiple lines, and dollar signs, please verify the operation thoroughly before using it. The behavior of these characters has not been fully verified. Check the test code for details.

- **Escaping in Environment Variables**:
  - **Double Quotes**: If a double quote (`"`) is present in the environment variable value, it will be escaped. For example, `double"quote` becomes `double\"quote`.
  - **New Lines**: If a new line character is present in the environment variable value, it will be escaped. For instance, a value `new
line` will be transformed to `new\nline`.

These transformations and precautions are important to maintain the integrity of data and the proper functioning of the system. Ensure that your input data is formatted and handled correctly to avoid any potential issues in processing.


# development
```
rm -rf bin/
go get github.com/mitchellh/gox
gox --output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"
```
