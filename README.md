aat-simplified
===

# Development

```BASH
git clone git@github.com:memochou1993/aat-simplified.git
git lfs clone git@github.com:memochou1993/aat.git
cp aat/storage/vocabulary.xml aat-simplified/vocabulary.xml
cd aat-simplified
go run main.go -f vocabulary.xml
./tb -f vocabulary.yaml
```
