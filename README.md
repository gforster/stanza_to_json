# stanza_to_json
Convert IBM stanza-style files to json and publish as an API

### Background
[IBM Stanza files](https://www.ibm.com/docs/en/spectrum-scale/5.0.0?topic=principles-stanza-files) have been extended to be used in other ways, but are very tied to AIX and legacy system administration.

I needed a way to eport valuable data from these stanza files to begin modernizing their use while not breaking current working production usage. This is meant to be an interim step to handling the data in these files in a modern approach.
