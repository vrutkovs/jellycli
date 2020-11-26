/*
 * Copyright 2020 Tero Vierimaa
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package api contains interface for connecting to remote server. Subpackages contain implementations.
package api

import (
	"tryffel.net/go/jellycli/interfaces"
	"tryffel.net/go/jellycli/models"
)

type ServerInfo struct {
	Name                 string
	Version              string
	Message              string
	RemoteControlEnabled bool
}

// Browser implements item-based viewing for music artists,albums,playlists etc.
type Browser interface {

	// GetArtists returns all artists
	GetArtists(paging interfaces.Paging) ([]*models.Artist, int, error)

	// GetAlbumArtists returns artists that are marked as album artists. See GetArtists.
	GetAlbumArtists(paging interfaces.Paging) ([]*models.Artist, int, error)
	// GetAlbums gets albums with given paging. Only PageSize and CurrentPage are used. Total count is returned
	GetAlbums(paging interfaces.Paging) ([]*models.Album, int, error)

	// GetArtistAlbums returns albums that artist takes part in.
	GetArtistAlbums(artist models.Id) ([]*models.Album, error)

	// GetAlbumSongs returns songs for given album id.
	GetAlbumSongs(album models.Id) ([]*models.Song, error)
	// GetPlaylists returns all playlists.
	GetPlaylists() ([]*models.Playlist, error)
	// GetPlaylistSongs fills songs array for playlist. If there's error, songs will not be filled
	GetPlaylistSongs(playlist models.Id) ([]*models.Song, error)
	// GetFavoriteArtists returns list of favorite artists.
	GetFavoriteArtists() ([]*models.Artist, error)
	// GetFavoriteAlbums return list of favorite albums.
	GetFavoriteAlbums(paging interfaces.Paging) ([]*models.Album, int, error)

	// GetSimilarArtists returns similar artists for artist id
	GetSimilarArtists(artist models.Id) ([]*models.Artist, error)

	// GetsimilarAlbums returns list of similar albums.
	GetSimilarAlbums(album models.Id) ([]*models.Album, error)

	// GetLatestAlbums returns latest albums.
	GetLatestAlbums() ([]*models.Album, error)

	// GetRecentlyPlayed returns songs that have been played last.
	GetRecentlyPlayed(paging interfaces.Paging) ([]*models.Song, int, error)

	// GetStatistics returns application statistics
	GetServerInfo() ServerInfo

	// GetSongs returns songs by paging. It also returns total number of songs.
	GetSongs(page, pageSize int) ([]*models.Song, int, error)

	// GetGenres returns music genres with paging. Return genres, total genres and possible error
	GetGenres(paging interfaces.Paging) ([]*models.IdName, int, error)

	// GetGenreAlbums returns all albums that belong to given genre
	GetGenreAlbums(genre models.IdName) ([]*models.Album, error)

	// GetAlbumArtist returns main artist for album.
	GetAlbumArtist(album *models.Album) (*models.Artist, error)

	// GetSongArtistAlbum returns main artist for song.
	GetSongArtistAlbum(song *models.Song) (*models.Album, *models.Artist, error)

	// GetInstantMix returns instant mix based on given item.
	GetInstantMix(item models.Item) ([]*models.Song, error)

	// GetLink returns a link to item that can be opened with browser.
	// If there is no link or item is invalid, empty link is returned.
	GetLink(item models.Item) string

	// Search returns values matching query and itemType, limited by number of maxResults,
	// Only items of itemType should ne returned.
	Search(query string, itemType models.ItemType, maxResults int) ([]models.Item, error)
}

// RemoteController implents controlling audio player remotely as well as
// keeping remote server updated on player status.
type RemoteController interface {
	// SetPlayer allows connecting remote controller to player, which can
	// then be controlled remotely.
	SetPlayer(player interfaces.Player)

	// ReportProgress reports player progress to remote controller.
	ReportProgress(state *interfaces.ApiPlaybackState) error
}
