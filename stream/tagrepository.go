package stream

import "sync"

// TagRepository manages sets of tags by key 
type TagRepository struct {
    mu sync.RWMutex
    tagsByKeyStore              map[string][]string
    keysByTagStore              map[string]map[string]bool
}

// NewTagRepository creates an empty tag repository
func NewTagRepository() TagRepository {
    return TagRepository {
        mu: sync.RWMutex {},
        tagsByKeyStore: make(map[string][]string),
        keysByTagStore: make(map[string]map[string]bool)}
}

// isStringInSlice returns true if the stringToFind exists within the strings slice
func isStringInSlice(strings []string, stringToFind string) bool {
    var matched = false
    
    if strings != nil {
        for _,v := range strings {
            if v == stringToFind {
                matched = true
                break      
            }
        }
    }
    
    return matched
}

// ApplyTags applies a set of tags to a key
func (tagRepository *TagRepository) ApplyTags(key string, tags []string, removeExisting bool) {
    defer tagRepository.mu.Unlock()
    tagRepository.mu.Lock()
    
    // get the current set of tags for a key
    var tagsForKey = tagRepository.tagsByKeyStore[key]
    
    // or make a new set if required
    if tagsForKey == nil || removeExisting {
        tagsForKey = make([]string, 100) 
    }
    
    // iterate through the tags to be applied
    for _,tag := range tags {
        if !isStringInSlice(tagsForKey, tag) {
            tagsForKey = append(tagsForKey, tag)
        } else {
            tagRepository.keysByTagStore[tag][key] = true
        }
    }
    
    tagRepository.tagsByKeyStore[key] = tagsForKey
}

// RemoveTags removes a set of tags from a key
func (tagRepository *TagRepository) RemoveTags(key string, tags []string) {
    defer tagRepository.mu.Unlock()
    tagRepository.mu.Lock()
    
    // get the current set of tags for a key
    var tagsForKey = tagRepository.tagsByKeyStore[key]
    
    if tagsForKey != nil {
        var newTags = make([]string, 10)
        
        for _,v := range tags {
            if !isStringInSlice(tags, v){
                newTags = append(newTags, v)
            } else {
                delete(tagRepository.keysByTagStore[v], key)
            }
        }
        
        tagRepository.tagsByKeyStore[key] = newTags
    }
}

// ClearForKey removes all tag information for akey
func (tagRepository *TagRepository) ClearForKey(key string) {
    defer tagRepository.mu.Unlock()
    tagRepository.mu.Lock()
    
    if tags,ok := tagRepository.tagsByKeyStore[key]; ok {
        if tags != nil {
            for _,tag := range tags {
                delete(tagRepository.keysByTagStore[tag], key)
            }
        }
        delete(tagRepository.tagsByKeyStore, key)   
    }
}

// Clear removes all tag information
func (tagRepository *TagRepository) Clear() {
    defer tagRepository.mu.Unlock()
    tagRepository.mu.Lock()
    
    tagRepository.keysByTagStore = make(map[string]map[string]bool,100)
    tagRepository.tagsByKeyStore = make(map[string][]string,100)    
}

// GetMatchingKeys returns all keys that have all tags within the withTags slice and no tags within the excludingTags slice
func (tagRepository *TagRepository) GetMatchingKeys(withTags []string, excludingTags []string) []string {
    defer tagRepository.mu.RUnlock()
    tagRepository.mu.RLock()
    
    return tagRepository.getMatchingKeysInternal(withTags, excludingTags)
}

// getMatchingKeysInternal returns all keys that have all tags within the withTags slice and no tags within the excludingTags slice
func (tagRepository *TagRepository) getMatchingKeysInternal(withTags []string, excludingTags []string) []string {
    var keysWithTagCount = make(map[string]int, 100)
    var excludedKeys = make(map[string]bool, 100)
    var expectedCount = len(withTags)
    var matchedKeys = make([]string,100)
    
    if excludingTags != nil {
        var keysToExclude = tagRepository.getMatchingKeysInternal(excludingTags, nil)
        
        for _,k := range keysToExclude {
            excludedKeys[k] = true
        }
    }
    
    if withTags != nil {
        for _,tag := range withTags {
            var keysWithTag = tagRepository.keysByTagStore[tag]
            
            if keysWithTag != nil {
                for key := range keysWithTag {
                    if currentCount, ok := keysWithTagCount[key]; ok {
                        keysWithTagCount[key] = currentCount + 1
                    } else{
                        keysWithTagCount[key] = 1
                    }
                }
            }
        }
        
        for k,v := range keysWithTagCount {
            if v == expectedCount {
                if _, inMap := excludedKeys[k]; !inMap {
                    matchedKeys = append(matchedKeys, k)
                }
            }
        }
    } else {
        // include all known keys that are not excluded
        for k := range tagRepository.tagsByKeyStore {
            if _, inMap := excludedKeys[k]; !inMap {
                matchedKeys = append(matchedKeys, k)
            }
        }
    }
    
    return matchedKeys
}

// GetTagsForKey returns a set of tags for a key
func (tagRepository *TagRepository) GetTagsForKey(key string) []string {
    defer tagRepository.mu.RUnlock()
    tagRepository.mu.RLock()
    
    return tagRepository.tagsByKeyStore[key]
}
