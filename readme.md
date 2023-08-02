ignite scaffold type batch proof:string batchnumber:uint
ignite scaffold map executionlayers name info chainid
ignite scaffold message addbatch chainid batch
ignite scaffold query getbatch chainid batchnumber --response proof:string
---
mapofmapd tx mapofmap addbatch chainid1234 sampleproof --from alice
mapofmapd query mapofmap getbatch chainid1234 1