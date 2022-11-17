# Instructor's notes

## Generate a CSV from `events.json`

```
jq -r '["timestamp", "version", "query_time"], (.[] | [.timestamp, .version, .query_time]) | @csv ' events.json
```
```csv
"timestamp","version","query_time"
1536051600,"356a192b7913b04c54574d18c28d46e6395428ab",189
1536832800,"77de68daecd823babbb58edb1c8e14d7106e83bb",124
1538704800,"902ba3cda1883801594b6e1b452790cc53948fda",210
1536202800,"da4b9237bacccdf19c0760cab7aec4a8359010b0",88
1536998400,"1b6453892473a467d07372d45eb05abc2031647a",202
1536728400,"77de68daecd823babbb58edb1c8e14d7106e83bb",313
1537531200,"1b6453892473a467d07372d45eb05abc2031647a",178
1538715600,"902ba3cda1883801594b6e1b452790cc53948fda",334
1536796800,"77de68daecd823babbb58edb1c8e14d7106e83bb",233
...
```

> See events.csv for complete list.

## Use a spreadsheet to calculate avg, min, and max query times

release|avg|min|max
---|---|---|---
77de68daecd823babbb58edb1c8e14d7106e83bb|201.83|88|336
902ba3cda1883801594b6e1b452790cc53948fda|208.93|88|334
c1dfd96eea8cc2b62785275bca38ac261256e278|211.33|91|304
ac3478d69a3c81fa62e60f5c3696165a4e5e6ac4|212.46|86|340
1b6453892473a467d07372d45eb05abc2031647a|214.44|86|339
356a192b7913b04c54574d18c28d46e6395428ab|215.32|86|341
da4b9237bacccdf19c0760cab7aec4a8359010b0|221.75|86|339
fe5dbbcea5ce7e2988b8c69bcfdfde8904aabc1f|224.71|89|340
0ade7c2cf97f75d009975f4d720d1fa6c19f4897|237.35|100|336

## Updated tests to use correct numbers

`TestAnalyzer_GetReleaseStats` and `TestAnalyzer_GetReleaseQuality` needed changes. See PR for diff.

After corrections, tests pass.

```
go test -v ./challenges/question1 
=== RUN   TestAnalyzer_GetReleaseStats
--- PASS: TestAnalyzer_GetReleaseStats (0.00s)
=== RUN   TestAnalyzer_GetReleaseQuality
--- PASS: TestAnalyzer_GetReleaseQuality (0.00s)
=== RUN   TestAnalyzer_GetReleaseHistory
--- PASS: TestAnalyzer_GetReleaseHistory (0.00s)
PASS
ok      github.com/jboursiquot/pivottechschool/challenges/question1     0.444s
```