

# safe reading in json to avoid shell expansion
```bash

# use a file first
jq -r '.value' filename


# read using echo or <<<
myJSON='{"aaaa": "bbb \"bb\" bb"}'
jq -r '.aaaa' <<<$myJSON
echo "$myJSON" | jq -r '.aaaa'
echo "$myJSON" | jq '.aaaa'

```

# iter with raw json
```bash
while read -r i; do
    # do stuff with $i
done < <(jq -c '.[]' <<<${json})
```


# multi-thread
```bash
function myFunc() {
    echo "hi from $1"
    sleep 5
    echo "bye from $1"
}

max_workers=10
workers=0
for i in $(seq 1 20); do
    [ $workers -eq $max_workers ] && {
        echo "Waiting for workers"
        wait
        workers=0
    }
    myFunc "$i" &
    workers=$((workers + 1))
done
wait
```