# #!/usr/bin/env bash
# set -e

# ORIGEN_PATH="out/out-clone-repos.json"
# URLS=$(cat $ORIGEN_PATH | grep "Error on to update repository - URL" | cut -d ' ' -f 8)
# DESTINATION_PATH="../out/out-run-clone-failured-tries.json"

# cd repos

# echo "" > $DESTINATION_PATH

# git config --global core.longpaths true

# git config --global http.version HTTP/1.1
# git config --global http.postBuffer 524288000
# git config --global http.lowSpeedLimit 0
# git config --global http.lowSpeedTime 999999
# git config --global core.protectNTFS false

# for item in $URLS
# do
#   echo $item >> $DESTINATION_PATH 2>> $DESTINATION_PATH
# done

# echo "#####################################" >> $DESTINATION_PATH 2>> $DESTINATION_PATH
# echo "#### Tring to clone repos" >> $DESTINATION_PATH 2>> $DESTINATION_PATH
# echo "#####################################" >> $DESTINATION_PATH 2>> $DESTINATION_PATH

# for item in $URLS
# do
#   git clone $item >> $DESTINATION_PATH 2>> $DESTINATION_PATH
# done

# echo "#####################################" >> $DESTINATION_PATH 2>> $DESTINATION_PATH
# echo "#### Tring to check pull repos" >> $DESTINATION_PATH 2>> $DESTINATION_PATH
# echo "#####################################" >> $DESTINATION_PATH 2>> $DESTINATION_PATH

# for item in $URLS
# do
#   splitedItems=(${item//\// })
#   repoName=${splitedItems[3]}
#   cd $repoName
#   echo $repoName
#   echo $item >> ../$DESTINATION_PATH 2>> ../$DESTINATION_PATH
#   git checkout -f >> ../$DESTINATION_PATH 2>> ../$DESTINATION_PATH
#   git pull >>  ../$DESTINATION_PATH 2>> ../$DESTINATION_PATH
#   cd ..
# done

# git config --global http.version HTTP/2
