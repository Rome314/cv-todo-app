echo "Running protoc..."
for D in ./rpc/*; do
        M=$(echo "$D" | cut -d "/" -f3)
        for F in ${D}/*.proto; do
                  mkdir ./api/"${M}";
                  protoc \
                      --proto_path=$GOPATH/src:. -I. -I $GOPATH/pkg/mod \
                      --twirp_out=. \
                      --go_out=. \
                      --doc_out=./api/"${M}" --doc_opt=markdown,README.md \
                      --twirp_swagger_out=./api/"${M}"/ \
                      --validate_out="lang=go:." \
                      "${F}"
        done
done

echo "Done."
