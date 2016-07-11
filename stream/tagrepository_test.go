package stream

import "testing"
import "github.com/stretchr/testify/assert"

// TestApplyAndGetTagsForKey tests repository calls for applying and retrieving tags by key
func TestApplyAndGetTagsForKey(t *testing.T) {
   var tagRepository = NewTagRepository()
   
   tagRepository.ApplyTags("myKey", []string{"First","Second"}, false)
   var tags = tagRepository.GetTagsForKey("myKey")
   
   assert.True(t, areSlicesEq(tags, []string{"First","Second"}, true))
   
   tagRepository.ApplyTags("myKey", []string{"Third"}, false)
   tags = tagRepository.GetTagsForKey("myKey")
   assert.True(t, areSlicesEq(tags, []string{"First","Second","Third"}, true))
   
   tagRepository.ApplyTags("myKey", []string{"Fourth","Fifth","Sixth"}, true)
   tags = tagRepository.GetTagsForKey("myKey")
   assert.True(t, areSlicesEq(tags, []string{"Fourth","Fifth","Sixth"}, true))
   
   tagRepository.ClearForKey("myKey")
   tags = tagRepository.GetTagsForKey("myKey")
   
   assert.True(t, areSlicesEq(tags, nil, true))
   tagRepository.ApplyTags("myKey", []string{"First","Second"}, false)
   
   tagRepository.Clear()
   tags = tagRepository.GetTagsForKey("myKey")
   assert.True(t, areSlicesEq(tags, nil, true))
}

// TestGetMatchingKeys tests repository calls for retrieving keys matching search conditions
func TestGetMatchingKeys(t *testing.T) {
   var tagRepository = NewTagRepository()
   
   tagRepository.ApplyTags("firstKey", []string{"A","B"}, false)
   tagRepository.ApplyTags("secondKey", []string{"A","C"}, false)
   tagRepository.ApplyTags("thirdKey", []string{"B","C"}, false)
   
   // test that only keys with the specified tag are returned
   var tags = tagRepository.GetMatchingKeys([]string{"B"}, nil)
   assert.True(t, areSlicesEq(tags, []string{"firstKey","thirdKey"}, false))
   
   // test that only keys with the specified tag AND not an exclude tag are returned
   tags = tagRepository.GetMatchingKeys([]string{"B"}, []string{"C"})
   assert.True(t, areSlicesEq(tags, []string{"firstKey"}, false))
   
   // test that no keys are returned when there are no keys with all tags
   tags = tagRepository.GetMatchingKeys([]string{"B","X"}, nil)
   assert.True(t, areSlicesEq(tags, []string{}, false))
   
   // test that all known keys are returned when there are no tags explicitly required
   tags = tagRepository.GetMatchingKeys(nil, nil)
   assert.True(t, areSlicesEq(tags, []string{"firstKey","secondKey","thirdKey"}, false))
   
   // test a call for all keys except those with a specified tag
   tags = tagRepository.GetMatchingKeys(nil, []string{"C"})
   assert.True(t, areSlicesEq(tags, []string{"firstKey"}, false))
}

// areSlicesEq determines if 2 slices are equal based on equality of their values, optionally in equal positions
func areSlicesEq(first, second []string, requireSamePositions bool) bool {
    var firstCount, secondCount = 0, 0
    
    if first == nil { 
        firstCount = 0 
    } else {
        firstCount = len(first)
    }
    
    if second == nil { 
        secondCount = 0 
    } else {
        secondCount = len(second)
    }
    
    if firstCount != secondCount {
        return false;
    }

    if firstCount > 0 { 
        var valueDict = make(map[string]int, firstCount)
        
        for i,v := range first {
            valueDict[v] = i
        }
        
        for i,v := range second {
            if pos,ok := valueDict[v]; ok {
                if requireSamePositions && i != pos {
                    return false; // value in wrong position
                }
            } else {
                return false; // value not found
            }
        } 
    }

    return true
}