package main

// Load attempts to load the swagger spec for the given package and given revision SHA.
// It returns the swagger spec content and true on success, nil and false if not found.
func Load(packagePath, sha string) ([]byte, bool) {
	return nil, false
}

// Save saves the given swagger spec to the cache.
func Save(b []byte, packagePath, sha string) {
}
