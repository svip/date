date.Date
=========

A date representations for Go.  Underlying is merely a `time.Time` object, and
it supports everything that `time.Time` can do.  It is function compatible with
`time.Time`, and can thus be used as a drop-in replacement.  Indeed, most of its
functions are merely call-throughs to `time.Time`'s functions.

[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/svip/date?tab=doc)

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

Reflections
-----------

It might seem questionable to create a library with a single `struct`, as it
feels it is approaching an almost `npm`-level of specialised packages.  And
indeed, if that concerns you, you can simply copy the code, as the Unlicence
provides no restrictions on the code.

However, there are times where the distinguishing between a date time and a
date is useful.  And using `time.Time` provides no such distinction.  Imagine
a library that builds a database based on `struct`s, the only way for it to do
so correctly, would be to use an annotation option.  But while that's available
for `struct`s, it won't be for function calls.

If there was a "standard" version of `date.Date`, that all these libraries
referred to, it would be easier to differentiate between date times and dates.
Though, perhaps a bit ambitious of me to suggest this might become any form of
"standard".

Another question to consider; why use `time.Time` as an underlying type, and
not simply just keep year, month and day?  And that might indeed a future
rewrite to avoid any potential problems with time zones and daylight saving
time, though I believe I have avoided the most of those faults, by enforcing the
`time.Time` instance to always be `time.UTC`.

But initially the motivation was simply that `time.Time` already had a lot of
functionality regarding dates, that on an initial version felt unnecessary to
re-implement.
