# FHERMA CIFAR-10 Challenge

## Introduction
[CIFAR-10](https://www.cs.toronto.edu/~kriz/cifar.html) is a widely recognized dataset comprising 60,000 color images of size 32x32 pixels, categorized into 10 classes (such as automobiles, airplanes, dogs, etc.). This dataset serves as a standard benchmark for machine learning algorithms in computer vision.

The goal of the challenge is to develop and implement a machine learning model capable of efficiently classifying encrypted CIFAR-10 images without decrypting them.


## Content
* `openfhe` - template for c++ openfhe based solution
* `openfhe-python` - template fo openfhe-python based solution
* `lattigo` - template for lattigo based solution
* `testcase.json` - testcase example(could be used with docker validator)

### How to validate solution locally
Once solution is developed a participants could use a docker image to validate their solution locally.
Put your solution with the test case json file in the local directory, link directory to the docker container and specify path to the project and test case json(keep in mind that path in arguments should be a path inside a docker container)

Example: local folder with the solution is `~/user/tmp/cifar10/app`
```sh
$ docker run -ti -v ~/user/tmp/cifar10:/cifar10 yashalabinc/fherma-validator --project-folder=/cifar10/app --testcase=/cifar10/testcase.json
```

Once validation is completed you'll see `result.json` file in the project folder. This file is exactly the same as we used to score uploaded solutions on the FHERMA platform
