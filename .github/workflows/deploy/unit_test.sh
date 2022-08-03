echo -e "------------ The script unit_test.sh is running ------------"
if [ -n "$1" ]; then
  noBaseDir="--noBaseDir"
  if [ "$1" != "$noBaseDir" ]; then
    cd $1 # The base directory
  fi
else
  echo "Missed the base directory param."
  exit 1
fi

# 1、Test the directory of apimock
echo "1. test the directory of apimock"
cd apimock && go test -v . && cd ../
if [ $? -ne 0 ]; then
  echo -e " >>>>>>>>>>>> Sorry, unit test failed <<<<<<<<<<<<"
  exit 1
fi

# 2、Test the directory of example/mock
echo "1. test the directory of example/mock"
cd example/mock && go test -v . && cd ../../
if [ $? -ne 0 ]; then
  echo -e " >>>>>>>>>>>> Sorry, unit test failed <<<<<<<<<<<<"
  exit 1
fi

echo -e "------------ The script unit_test.sh is stopped ------------"
