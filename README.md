This is command line tool for Machine Learning, which you can used to train ML models and deploy models as service/api.

## `rexcli --help`

```
github.com/RexG$ rexcli --help
A Machine Learning command line tool to help you train models and deploy models as service/api

Usage:
  rexcli [command]

Available Commands:
  help        Help about any command
  init        Prepare the Machine Learning environment on your local machine
  model       Train/deploy/undeploy/push/pull ML models
  model-api   Check your model-api status/log
  version     Print current version
```

## `rexcli model --help`
```
src/github.com/RexG$ rexcli model --help
Train/deploy/undeploy/push/pull ML models

Usage:
  rexcli model [flags]
  rexcli model [command]

Available Commands:
  deploy      Deploy your model in  as a model-api (Restful API)
  list        List all your models from model repository
  train       Train your ML model
  upload      Upload local model to model repository
```

### `rexcli model train --help`

> if it's local, then will create a **Jupyter** instance for you

```
github.com/RexG$ rexcli model train --help
Train your ML model

Usage:
  rexcli model train [flags]

Flags:
  -p, ---prod         Train ML models in  PROD Jupyter notebooks
  -s, ---stg          Train ML models in  STG Jupyter notebooks
  -h, --help          help for train
  -l, --local         Train ML models in your local Jupyter notebook
      --name string   Name your local Jupyter instance (default "local-")
      --port int      Specify your local Jupyter instance port (default 8888)
```

### `rexcli model deploy --help`

```
github.com/RexG$ rexcli model deploy --help
Deploy your model in  as a model-api (Restful API)

Usage:
  rexcli model deploy [flags]

Flags:
      ---prod                    deploy your model in  PROD
      ---stg                     deploy your model in  STG
      --api-name string          your model-api name
      --cpu int                  how many CPUs (default 1)
  -h, --help                     help for deploy
      --memory int               how many Gi memory for 1 model-api (default 1)
      --model-framework string   which ML framework you used fro training this model, e.g. MLflow/Tensorflow/XGBoost/SKLearn (default "MLflow")
      --model-url string         where your model is, a URL
      --replica-set int          how many model-api instances for 1 model-api (default 1)
```