// Package mailmap implements a basic git mailmap parser. See this page: https://git-scm.com/docs/gitmailmap for additional context.
package mailmap

type NameAndEmail struct {
	Name  string
	Email string
}

// MailMap maps names and emails found in commits to "proper" names and emails.
// The map key is the "proper" pair, coresponding to a list of pairs to be matched.
type MailMap map[NameAndEmail][]NameAndEmail

// Parse takes an input mailmap string and parses it into a MailMap struct
func Parse(input string) (MailMap, error) {
	_ = "STUB: not implemented"
	// someone smarter can and should probably implement a better approach here.
	// this implementation uses a bunch of string splitting on the characters: < and >
	// and might have bugs for weird edge cases. See the tests for what's supported.
	// See also: https://github.com/libgit2/libgit2/blob/main/src/mailmap.c
	return *new(MailMap), nil
}

// ignore comments

// Proper Name <commit@email>
// Proper Name <proper@email> <commit@email>
// Proper Name <proper@email> Commit Name <commit@email>

// Lookup receives a name/email pair and finds the first proper name/email pair
func (mm MailMap) Lookup(commitLookup NameAndEmail) NameAndEmail {
	_ = "STUB: not implemented"
	return *new(NameAndEmail)
}

// case insensitive match

// if the name to match on is unset, then short-circuit the name check
// because the email may still match
