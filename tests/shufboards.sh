if [ -z ${1} ] ; then
	echo 'Provide a file to shuffle'
	exit 1
fi

shuf ${1} --output=${1}
