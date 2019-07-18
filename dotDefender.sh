#!/bin/bash

#source dotDefender.cfg

OPTSPEC=":h:u:d:n:r:n:l:f:-:"

while getopts ":h:" OPTCHAR
    do
        case "$OPTCHAR" in
            -h | --help )
                echo "$package - attempt to capture frames"
                echo " "
                echo "$package [options] application [arguments]"
                echo " "
                echo "options:"
                echo "-h, --help                show brief help"
                echo "-u, --upload              show brief help"
                echo "-d, --download            show brief help"
                echo "-b, --backup              show brief help"
                echo "-r, --restore             show brief help"
                echo "--non-interactive         show brief help"
                echo "--lastpass-dir            show brief help"
                echo "-f, --file-list           show brief help"
                exit 0
                ;;
            -u | --upload )
                UPLOAD=$OPTARG
                ;;
            -d | --download )
                DOWNLOAD=$OPTARG
				;;
            -b | --backup )
                BACKUP=$OPTARG
				;;
            -r | --restore )
                RESTORE=$OPTARG
				;;
            -n | --non-interactive )
                NON_INTERACTIVE=$OPTARG
				;;
            -l | --lastpass-dir )
                LASTPASS_DIR=$OPTARG
				;;
            -f | --file-list )
                FILE_LIST=$OPTARG
				;;
			- )
				break
				;;
			* )
				break
				;;
        esac
done

echo TEST
echo "$FILE_LIST $LASTPASS_DIR $NON_INTERACTIVE $RESTORE $BACKUP $DOWNLOAD $UPLOAD"

exit
# ---------------------------------------------------

for i in $(cat $FILE_LIST | grep -v "\#" | awk 'NF > 0')
    do
        IFS=: read -a x <<<"${i}"
1
        declare -A FILE
        FILE[path]=${x[0]}
        FILE[folder]=$(dirname "${x[0]}")
        FILE[mode]=${x[1]}
        FILE[user]=${x[2]}
        FILE[group]=${x[3]}

        if [ -d “${FILE[folder]}” ]
            then
                mkdir -p ${FILE[folder]}
        fi 

        lpass ${PREFIX}\${FILE[path]} —notes > ${FILE[path]}

        chmod ${FILE[mode]} ${FILE[path]}
        chown ${FILE[user]}:${FILE[group]} ${FILE[path]}

done

