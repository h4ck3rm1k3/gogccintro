
curl -vvv http://localhost:9999/blazegraph/sparql -X POST --data 'query=SELECT+%2A+WHERE+%7B+%3Fs+%3Fp+%3Fo+%7D+LIMIT+1' -H 'Accept:application/rdf+xml'
