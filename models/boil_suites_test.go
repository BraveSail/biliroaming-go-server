// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("AccessKeys", testAccessKeys)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCaches)
	t.Run("PlayURLCaches", testPlayURLCaches)
	t.Run("SeasonAreaCaches", testSeasonAreaCaches)
	t.Run("THEpisodeCaches", testTHEpisodeCaches)
	t.Run("THSeason2Caches", testTHSeason2Caches)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCaches)
	t.Run("THSeasonCaches", testTHSeasonCaches)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCaches)
	t.Run("THSubtitleCaches", testTHSubtitleCaches)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("AccessKeys", testAccessKeysDelete)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesDelete)
	t.Run("PlayURLCaches", testPlayURLCachesDelete)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesDelete)
	t.Run("THEpisodeCaches", testTHEpisodeCachesDelete)
	t.Run("THSeason2Caches", testTHSeason2CachesDelete)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesDelete)
	t.Run("THSeasonCaches", testTHSeasonCachesDelete)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesDelete)
	t.Run("THSubtitleCaches", testTHSubtitleCachesDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("AccessKeys", testAccessKeysQueryDeleteAll)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesQueryDeleteAll)
	t.Run("PlayURLCaches", testPlayURLCachesQueryDeleteAll)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesQueryDeleteAll)
	t.Run("THEpisodeCaches", testTHEpisodeCachesQueryDeleteAll)
	t.Run("THSeason2Caches", testTHSeason2CachesQueryDeleteAll)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesQueryDeleteAll)
	t.Run("THSeasonCaches", testTHSeasonCachesQueryDeleteAll)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesQueryDeleteAll)
	t.Run("THSubtitleCaches", testTHSubtitleCachesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("AccessKeys", testAccessKeysSliceDeleteAll)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesSliceDeleteAll)
	t.Run("PlayURLCaches", testPlayURLCachesSliceDeleteAll)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesSliceDeleteAll)
	t.Run("THEpisodeCaches", testTHEpisodeCachesSliceDeleteAll)
	t.Run("THSeason2Caches", testTHSeason2CachesSliceDeleteAll)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesSliceDeleteAll)
	t.Run("THSeasonCaches", testTHSeasonCachesSliceDeleteAll)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesSliceDeleteAll)
	t.Run("THSubtitleCaches", testTHSubtitleCachesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("AccessKeys", testAccessKeysExists)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesExists)
	t.Run("PlayURLCaches", testPlayURLCachesExists)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesExists)
	t.Run("THEpisodeCaches", testTHEpisodeCachesExists)
	t.Run("THSeason2Caches", testTHSeason2CachesExists)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesExists)
	t.Run("THSeasonCaches", testTHSeasonCachesExists)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesExists)
	t.Run("THSubtitleCaches", testTHSubtitleCachesExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("AccessKeys", testAccessKeysFind)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesFind)
	t.Run("PlayURLCaches", testPlayURLCachesFind)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesFind)
	t.Run("THEpisodeCaches", testTHEpisodeCachesFind)
	t.Run("THSeason2Caches", testTHSeason2CachesFind)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesFind)
	t.Run("THSeasonCaches", testTHSeasonCachesFind)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesFind)
	t.Run("THSubtitleCaches", testTHSubtitleCachesFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("AccessKeys", testAccessKeysBind)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesBind)
	t.Run("PlayURLCaches", testPlayURLCachesBind)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesBind)
	t.Run("THEpisodeCaches", testTHEpisodeCachesBind)
	t.Run("THSeason2Caches", testTHSeason2CachesBind)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesBind)
	t.Run("THSeasonCaches", testTHSeasonCachesBind)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesBind)
	t.Run("THSubtitleCaches", testTHSubtitleCachesBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("AccessKeys", testAccessKeysOne)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesOne)
	t.Run("PlayURLCaches", testPlayURLCachesOne)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesOne)
	t.Run("THEpisodeCaches", testTHEpisodeCachesOne)
	t.Run("THSeason2Caches", testTHSeason2CachesOne)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesOne)
	t.Run("THSeasonCaches", testTHSeasonCachesOne)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesOne)
	t.Run("THSubtitleCaches", testTHSubtitleCachesOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("AccessKeys", testAccessKeysAll)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesAll)
	t.Run("PlayURLCaches", testPlayURLCachesAll)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesAll)
	t.Run("THEpisodeCaches", testTHEpisodeCachesAll)
	t.Run("THSeason2Caches", testTHSeason2CachesAll)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesAll)
	t.Run("THSeasonCaches", testTHSeasonCachesAll)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesAll)
	t.Run("THSubtitleCaches", testTHSubtitleCachesAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("AccessKeys", testAccessKeysCount)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesCount)
	t.Run("PlayURLCaches", testPlayURLCachesCount)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesCount)
	t.Run("THEpisodeCaches", testTHEpisodeCachesCount)
	t.Run("THSeason2Caches", testTHSeason2CachesCount)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesCount)
	t.Run("THSeasonCaches", testTHSeasonCachesCount)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesCount)
	t.Run("THSubtitleCaches", testTHSubtitleCachesCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("AccessKeys", testAccessKeysHooks)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesHooks)
	t.Run("PlayURLCaches", testPlayURLCachesHooks)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesHooks)
	t.Run("THEpisodeCaches", testTHEpisodeCachesHooks)
	t.Run("THSeason2Caches", testTHSeason2CachesHooks)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesHooks)
	t.Run("THSeasonCaches", testTHSeasonCachesHooks)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesHooks)
	t.Run("THSubtitleCaches", testTHSubtitleCachesHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("AccessKeys", testAccessKeysInsert)
	t.Run("AccessKeys", testAccessKeysInsertWhitelist)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesInsert)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesInsertWhitelist)
	t.Run("PlayURLCaches", testPlayURLCachesInsert)
	t.Run("PlayURLCaches", testPlayURLCachesInsertWhitelist)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesInsert)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesInsertWhitelist)
	t.Run("THEpisodeCaches", testTHEpisodeCachesInsert)
	t.Run("THEpisodeCaches", testTHEpisodeCachesInsertWhitelist)
	t.Run("THSeason2Caches", testTHSeason2CachesInsert)
	t.Run("THSeason2Caches", testTHSeason2CachesInsertWhitelist)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesInsert)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesInsertWhitelist)
	t.Run("THSeasonCaches", testTHSeasonCachesInsert)
	t.Run("THSeasonCaches", testTHSeasonCachesInsertWhitelist)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesInsert)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesInsertWhitelist)
	t.Run("THSubtitleCaches", testTHSubtitleCachesInsert)
	t.Run("THSubtitleCaches", testTHSubtitleCachesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("AccessKeyToUserUsingUIDUser", testAccessKeyToOneUserUsingUIDUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("UserToUIDAccessKeys", testUserToManyUIDAccessKeys)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("AccessKeyToUserUsingUIDAccessKeys", testAccessKeyToOneSetOpUserUsingUIDUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("UserToUIDAccessKeys", testUserToManyAddOpUIDAccessKeys)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("AccessKeys", testAccessKeysReload)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesReload)
	t.Run("PlayURLCaches", testPlayURLCachesReload)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesReload)
	t.Run("THEpisodeCaches", testTHEpisodeCachesReload)
	t.Run("THSeason2Caches", testTHSeason2CachesReload)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesReload)
	t.Run("THSeasonCaches", testTHSeasonCachesReload)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesReload)
	t.Run("THSubtitleCaches", testTHSubtitleCachesReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("AccessKeys", testAccessKeysReloadAll)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesReloadAll)
	t.Run("PlayURLCaches", testPlayURLCachesReloadAll)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesReloadAll)
	t.Run("THEpisodeCaches", testTHEpisodeCachesReloadAll)
	t.Run("THSeason2Caches", testTHSeason2CachesReloadAll)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesReloadAll)
	t.Run("THSeasonCaches", testTHSeasonCachesReloadAll)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesReloadAll)
	t.Run("THSubtitleCaches", testTHSubtitleCachesReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("AccessKeys", testAccessKeysSelect)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesSelect)
	t.Run("PlayURLCaches", testPlayURLCachesSelect)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesSelect)
	t.Run("THEpisodeCaches", testTHEpisodeCachesSelect)
	t.Run("THSeason2Caches", testTHSeason2CachesSelect)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesSelect)
	t.Run("THSeasonCaches", testTHSeasonCachesSelect)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesSelect)
	t.Run("THSubtitleCaches", testTHSubtitleCachesSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("AccessKeys", testAccessKeysUpdate)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesUpdate)
	t.Run("PlayURLCaches", testPlayURLCachesUpdate)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesUpdate)
	t.Run("THEpisodeCaches", testTHEpisodeCachesUpdate)
	t.Run("THSeason2Caches", testTHSeason2CachesUpdate)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesUpdate)
	t.Run("THSeasonCaches", testTHSeasonCachesUpdate)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesUpdate)
	t.Run("THSubtitleCaches", testTHSubtitleCachesUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("AccessKeys", testAccessKeysSliceUpdateAll)
	t.Run("EpisodeAreaCaches", testEpisodeAreaCachesSliceUpdateAll)
	t.Run("PlayURLCaches", testPlayURLCachesSliceUpdateAll)
	t.Run("SeasonAreaCaches", testSeasonAreaCachesSliceUpdateAll)
	t.Run("THEpisodeCaches", testTHEpisodeCachesSliceUpdateAll)
	t.Run("THSeason2Caches", testTHSeason2CachesSliceUpdateAll)
	t.Run("THSeason2EpisodeCaches", testTHSeason2EpisodeCachesSliceUpdateAll)
	t.Run("THSeasonCaches", testTHSeasonCachesSliceUpdateAll)
	t.Run("THSeasonEpisodeCaches", testTHSeasonEpisodeCachesSliceUpdateAll)
	t.Run("THSubtitleCaches", testTHSubtitleCachesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
