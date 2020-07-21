5 * * * * go run sh


sed -i "s/banana/apple/g" user.txt

vmstat 2 5


split -l 3 test  file




line=`wc -l test`
k=1
for i in `seq 1 $line`
do
	sed -n $kp test >> res
	if [ i%3 == 0]
		k = 1
	fi
done
