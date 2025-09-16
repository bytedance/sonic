# JSON Test Suite

["Parsing JSON is a Minefield 💣"](https://seriot.ch/projects/parsing_json.html)
(posted 2016-10-26) performed one of the first thorough comparisons of
JSON parser implementations and their behavior on various edge-cases.
The test cases from that article have been copied into this directory.
At the time, [RFC 7159](https://www.rfc-editor.org/rfc/rfc7159.html)
was the authoritative standard, but has since been superseded by
[RFC 8259](https://www.rfc-editor.org/rfc/rfc8259.html).
Consequently, the expected results of some of the test cases from the article
were changed to be more compliant with RFC 8259.

# Changes to test cases

## Require rejection of invalid UTF-8

[RFC 8259, section 8.1](https://www.rfc-editor.org/rfc/rfc8259.html#section-8.1)
requires that JSON text be formatted using UTF-8.

The classification of the following cases was changed:

| Case name                             | Verdict difference               |
| ------------------------------------- | -------------------------------- |
| string_invalid_utf-8                  | either pass or fail ⇨ must fail |
| string_UTF8_surrogate_U+D800          | either pass or fail ⇨ must fail |
| string_UTF-8_invalid_sequence         | either pass or fail ⇨ must fail |
| string_iso_latin_1                    | either pass or fail ⇨ must fail |
| string_lone_utf8_continuation_byte    | either pass or fail ⇨ must fail |
| string_not_in_unicode_range           | either pass or fail ⇨ must fail |
| string_overlong_sequence_2_bytes      | either pass or fail ⇨ must fail |
| string_overlong_sequence_6_bytes      | either pass or fail ⇨ must fail |
| string_overlong_sequence_6_bytes_null | either pass or fail ⇨ must fail |
| string_truncated-utf-8                | either pass or fail ⇨ must fail |
| string_UTF-16LE_with_BOM              | either pass or fail ⇨ must fail |
| string_utf16BE_no_BOM                 | either pass or fail ⇨ must fail |
| string_utf16LE_no_BOM                 | either pass or fail ⇨ must fail |

One exception is that a byte order mark (U+FEFF) at the start may be ignored
by an implementation (in contrast to treating it as an error). For this reason,
"structure_UTF-8_BOM_empty_object" is left as "either pass or fail".

[RFC 8259, section 8.2](https://www.rfc-editor.org/rfc/rfc8259.html#section-8.2)
specifies that it is undefined how invalid escaped surrogate pairs are handled.
An implementation may accept or reject such cases.

The classification of the following cases was left unchanged:

| Case name                                    | Verdict             |
| -------------------------------------------- | ------------------- |
| structure_UTF-8_BOM_empty_object             | either pass or fail |
| object_key_lone_2nd_surrogate                | either pass or fail |
| string_1st_surrogate_but_2nd_missing         | either pass or fail |
| string_1st_valid_surrogate_2nd_invalid       | either pass or fail |
| string_incomplete_surrogate_and_escape_valid | either pass or fail |
| string_incomplete_surrogate_pair             | either pass or fail |
| string_incomplete_surrogates_escape_valid    | either pass or fail |
| string_invalid_lonely_surrogate              | either pass or fail |
| string_invalid_surrogate                     | either pass or fail |
| string_inverted_surrogates_U+1D11E           | either pass or fail |
| string_lone_second_surrogate                 | either pass or fail |

Note that these cases are expected to be rejected under
[RFC 7493, section 2.1](https://datatracker.ietf.org/doc/html/rfc7493#section-2.1).
RFC 7493 is compatible with RFC 8259 in that it makes strict decisions
about behavior that RFC 8259 leaves undefined.

## Permit rejection of duplicate object names

[RFC 8259, section 4](https://datatracker.ietf.org/doc/html/rfc8259#section-4)
says:

> When the names within an object are not unique,
> the behavior of software that receives such an object is unpredictable.
> Many implementations report the last name/value pair only.
> Other implementations report an error or fail to parse the object, and
> some implementations report all of the name/value pairs, including duplicates.

Thus, handling of duplicate object names is undefined behavior.
Rejecting such occurences is within the realm of permitted behavior.

The classification of the following cases was changed:

| Case name                       | Verdict difference               |
| ------------------------------- | -------------------------------- |
| object_duplicated_key_and_value | must pass ⇨ either pass or fail |
| object_duplicated_key           | must pass ⇨ either pass or fail |

Note that these cases are expected to be rejected under
[RFC 7493, section 2.3](https://datatracker.ietf.org/doc/html/rfc7493#section-2.3):

> Objects in I-JSON messages MUST NOT have members with duplicate names.
> In this context, "duplicate" means that the names,
> after processing any escaped characters,
> are identical sequences of Unicode characters.

RFC 7493 is compatible with RFC 8259 in that it makes strict decisions
about behavior that RFC 8259 leaves undefined.

This decision is further guided by a number of real security vulnerabilities
that relied on duplicate object names to bypass checks:
* https://justi.cz/security/2017/11/14/couchdb-rce-npm.html
* https://nvd.nist.gov/vuln/detail/cve-2022-25757
* https://bishopfox.com/blog/json-interoperability-vulnerabilities

## Require acceptance of large numbers

[RFC 8259, section 6](https://www.rfc-editor.org/rfc/rfc8259.html#section-6)
describes the ABNF grammar for a JSON number that can have arbitrarily
large representations. It does warn that implementations may not be
able to represent the JSON number. However, the expected failure mode seems
to be one where the implementation "will approximate JSON numbers within
the expected precision" rather than outright fail on parsing.

[RFC 8259, section 9](https://www.rfc-editor.org/rfc/rfc8259.html#section-9)
later says that an "implementation may set limits on the range and precision
of numbers." However, this is in the context transforming JSON text into
some other data representation. Our tests are only concerned about whether
we can validate the input JSON, and not about transformation.
It's a question of the difference between syntax and semantics.
Thus, this exemption clause does not apply in our context.

The classification of the following cases was changed:

| Case name                    | Verdict difference               |
| ---------------------------- | -------------------------------- |
| number_double_huge_neg_exp   | either pass or fail ⇨ must pass |
| number_huge_exp              | either pass or fail ⇨ must pass |
| number_neg_int_huge_exp      | either pass or fail ⇨ must pass |
| number_pos_double_huge_exp   | either pass or fail ⇨ must pass |
| number_real_neg_overflow     | either pass or fail ⇨ must pass |
| number_real_pos_overflow     | either pass or fail ⇨ must pass |
| number_real_underflow        | either pass or fail ⇨ must pass |
| number_too_big_neg_int       | either pass or fail ⇨ must pass |
| number_too_big_pos_int       | either pass or fail ⇨ must pass |
| number_very_big_negative_int | either pass or fail ⇨ must pass |
