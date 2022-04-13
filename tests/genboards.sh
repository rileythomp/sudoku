if [ -z ${1} ] ; then
	echo 'Provide a difficulty as first arg'
	exit 1
fi

if [ -x ${2} ]; then
	echo 'Provide a # of boards as second arg'
	exit 1
fi

for i in $(seq ${2}); do
	curl --silent https://sugoku2.herokuapp.com/board?difficulty=${1} | sed 's/[^0-9]*//g' | grep -v Total
done
