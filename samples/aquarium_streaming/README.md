# RPI3 Aquarium Streaming
- It consists of RPI3, Pi Cam, AWS IoT Core, AWS Lamda Function (C2C Connecotr), SmartThings DevWorkspace and SmartThings App

## HOWTO
### 1. Register a thing in AWS IoT Core
- Refer guide from AWS Documentation
- https://docs.aws.amazon.com/en_pv/iot/latest/developerguide/register-device
![](/media/register_device_in_aws_iot_core.png)

- Or use full PDF guide
- https://docs.aws.amazon.com/en_pv/iot/latest/developerguide/iot-dg.pdf#register-device

### 2. Run AWS IoT Client in a thing
- Clone AWS IoT Core SDK from GitHub
- https://github.com/aws/aws-iot-device-sdk-python.git
![](/media/aws_iot_sdk_python_github.png)

- Run sample of ThingShadowEcho.py with parameters prepared by No.1
```shell
$ python ./ThingShadowEcho.py  -e YOUR_AWS_IOT_CUSTOM_ENDPOINT_URL -r YOU_ROOT_CA_FILE_PATH -c YOUR_CERTIFICATE_FILE_PATH -k YOUR_PRIVATE_KEY_FILE_PATH -n YOUR_TARGETED_THING_NAME
```

### 3. Register a C2C connector in AWS Lamda Function and SmartThings DeveloperWorkspace
- Refer guide from SmartThings Documentation
- https://smartthings.developer.samsung.com/docs/devices/smartthings-schema/schema-basics.html#Cloud-Connector
![](/media/cloud_connector_with_st_devws.png)

- https://smartthings.developer.samsung.com/docs/devices/smartthings-schema/schema-basics.html#Register-Connector-in-Developer-Workspace
![](/media/register_connector_with_st_devws.png)

### 4. Enable Devloper Mode in SmartThings App 
- Refer guide from SmartThings Documentation
- https://smartthings.developer.samsung.com/docs/testing/how-to-test.html
![](/media/test_device_with_smartthings_app.png)

