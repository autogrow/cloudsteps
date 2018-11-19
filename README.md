# Cloud Steps

A tool to merge a YAML step function definition in with a CloudFormation template file.

## Usage

You have this at `steps.yml`:

```yml
StateMachineName: HelloWorld-StateMachine
DefinitionString:
    StartAt: HelloWorld
    States:
    HelloWorld:
        Type:     Task
        Resource: "arn:aws:lambda:${AWS::Region}:${AWS::AccountId}:function:HelloFunction"
        End:      true
RoleArn: !Sub "arn:aws:iam::${AWS::AccountId}:role/service-role/StatesExecutionRole-${AWS::Region}"
```

But you want to put it in your cloud formation template but you [came across this problem](https://stackoverflow.com/questions/51627531/deploy-stepfunctions-with-cloudformation-from-external-definition-file).  Run the command like so:

    cloudsteps -t cftemplate.yml -in steps.yml -t

And it will be added to the resources in the template but embedded as JSON:

```yml
AWSTemplateFormatVersion: '2010-09-09'
Description: An example template for a Step Functions state machine.
Resources:
  HelloWorld-StateMachine:
    Type: AWS::StepFunctions::StateMachine
    Properties:
      StateMachineName: HelloWorld-StateMachine
      DefinitionString:
        Fn::Sub: '{"HelloWorld":{"End":true,"Resource":"arn:aws:lambda:${AWS::Region}:${AWS::AccountId}:function:HelloFunction","Type":"Task"},"StartAt":"HelloWorld","States":null}'
      RoleArn: !Sub arn:aws:iam::${AWS::AccountId}:role/service-role/StatesExecutionRole-${AWS::Region}
```