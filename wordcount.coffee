wordcount = (s) ->
  words = {}
  words[word] = words[word] + 1 || 1 for word in s.split(/\s+/)
  ([key, value] for key, value of words).sort (a,b) ->
    if a[1] > b[1] then -1 else 1
