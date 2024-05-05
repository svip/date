date.Date
=========

A date representations for Go.  Underlying is merely a `time.Time` object, and
it supports everything that `time.Time` can do.  It is function compatible with
`time.Time`, and can thus be used as a drop-in replacement.  Indeed, most of its
functions are merely call-throughs to `time.Time`'s functions.

Motivations
-----------

Consider the following JSON object using ISO 8601 date format:

```
{
	"date": "2024-06-05"
}
```

If you wished to use `encoding/json` to unmarshal that JSON into a Go `struct`,
making your date representation a `time.Time` would not work:

```
type Data struct {
	Date time.Time `json:"date"`
}
```

`json.Unmarshal` would only yield the following error:

```
parsing time "2024-06-05" as "2006-01-02T15:04:05Z07:00": cannot parse "" as "T"
```

This is because `time.Time.UnmarshalJSON` assumes the RFC 3339 date time format,
and will not use the ISO 8601 date format as a fallback.

Commonly, the solution is to simply use a `string` here, and then use
`time.Parse()` later with the ISO 8601 date representation.  But that can be a
little bit tedious.  So using this library, you can simply use `date.Date` for
the `Date`-field, and it'll parse the ISO 8601 date representation.

```
type Data struct {
	Date date.Date `json:"date"`
}
```

The `date.Date` has a `.Time()` function that returns the underlying `time.Time`
representation.

Considerations
--------------

`date.Date` is not built for speed, and its parsing could be better optimised,
so you may wish to avoid it if performance is of the essence in those instances.
Though, in those situations, perhaps even using `time.Time` is too slow.

Its `GobDecode`, `GobEncode`, `MarshalBinary` and `UnmarshalBinary` uses the
same binary representation as `time.Time`, but the time information will be
lost.

It's also worth considering whether another dependency is worth it, to which
I suggest only the implementation of `UnmarshalJSON` and `Time() time.Time` are
truly necessary.

Interface compatible
--------------------

As mentioned, it has all the same functions as `time.Time`, with the
`time.Time` type for variables having been replaced with `date.Date`, in all
but one case.  This means you can use it as a drop-in replacement, but when
comparing to other times, they must be `date.Date`.
