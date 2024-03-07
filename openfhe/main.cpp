#include <iostream>
#include <fstream>
#include <string>
#include "cifar10.h"

using namespace lbcrypto;

int main(int argc, char *argv[])
{
    std::string pubKeyLocation;
    std::string multKeyLocation;
    std::string rotKeyLocation;
    std::string ccLocation;
    std::string inputLocation;
    std::string outputLocation;

    for (auto i = 1; i < argc; i += 2)
    {
        std::string arg = argv[i];
        if (arg == "--key_pub")
        {
            pubKeyLocation = argv[i + 1];
        }
        else if (arg == "--key_mult")
        {
            multKeyLocation = argv[i + 1];
        }
        else if (arg == "--key_rot")
        {
            rotKeyLocation = argv[i + 1];
        }
        else if (arg == "--cc")
        {
            ccLocation = argv[i + 1];
        }
        else if (arg == "--input")
        {
            inputLocation = argv[i + 1];
        }
        else if (arg == "--output")
        {
            outputLocation = argv[i + 1];
        }
    }

    CIFAR10CKKS CIFAR10CKKS(ccLocation, pubKeyLocation, multKeyLocation, rotKeyLocation, inputLocation, 
                                  outputLocation);
    CIFAR10CKKS.eval();
    CIFAR10CKKS.serializeOutput();
    return 0;
}
