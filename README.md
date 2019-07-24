# Aerospike-Elexir-Port

Golang Prog for Uniris Node to act as intermediate between the erlang VM and an aerospike instance

## Examples

### Get request

{"QueryID" : "123", "Type" : "get", "Namespace" : "on-disk-db", "Set" : "test", "Key" : "7157762ecb437e34e5770667d086887532a304d9152b171a974d100f2fcfb687"}

### Put request

{"QueryID" : "123", "Type" : "put", "Namespace" : "on-disk-db", "Set" : "test", "Key" : "7157762ecb437e34e5770667d086887532a304d9152b171a974d100f2fcfb687", "Bins" : [{ "BinName" : "addr", "BinValue" :"7157762ecb437e34e5770667d086887532a304d9152b171a974d100f2fcfb687"}, {"BinName" : "data","BinValue" :"JUbqK5m6c7I4zOKZiox4uLfmyJ5Qrj6Fbm4OZQeJdMNOFRZOVBr00EHCQ7ld0l7edTVSl0Bw6OfK7jV5eltq7nnyVXHIsK5pjyEs4W42kUQ6yIlNdjIF9GDxhQTuuO5Mn3MDEOtO6XnNqj2GXmt2NgNWD2JIBWxbdzhw5qD4ifQQrnNYQ3UG63eITM8AQD1lHONimYqI41tAZoFWHafqORaGpfO94rO1FKt3pRYnH9Nv9J0YG5wCm9wN2cXfdSbgibRGK4AlK7MmurMopnewDjNdGx0ecZHPxFQdpb6G0v8w1pCEZzZhDI5SltZPVippLfPXGsKA03vNoMEtKldObcEZE5bbXVaovi63doHzxpx9taJ3NY8km2kCRaXte8FBiONrIwNMcXRfto08nRcEICmFuHUTb9mG7GqwCmbBMc8EP3GJG5vUXgJnHF2uLpcuA4BiRNYPeuHJUXGQIxRY3CNQj5ew4fDqsHnNo2lfeUA0Qgxobks8ElDb9RnC1FcqpwsYcAGxpnB0EeQxuqzAOaiWR2pnzpreQXAlaroF8Sdkc25vIGtNaeHyOv7Q62P2IHrjLXMiFYGBtpA0ahiB1BJru7KmgEgwJ5vHseTjw7gN5szETSj2iPlv0SjgFSjlH5zUETlZQw0EFJVQOuqbg89BZAFYv8sbDnyw1aoP3R66o0hTvzIuGiSW6c5m4CSFyc8doeN91u7FYzWttSRY56Hdf53QDF5ax2KaBgk396Ni45WZu1q70wbxolmNVXn2YrMeWNW7cSMXPNRzVCGR3BSQvF51EPSBko8OJ1cif5deLfGbM3aJsWPNISuC3jAWtIT9Fdfrc7PVdCiVcDMj0b8IgMfw1qupvxBIV27AU02nRqADoFX0Fc2M09EAUEXUEriLpbUtLC3sUeIykAlC1OPv9WJxcErShxWeXDBXhNh1mQDgWMOHXqh0yWODAPTOq27uAbdWwAFjiYYigCAn5OGwOyakHu7Xc93fLHVc" }]}