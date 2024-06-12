// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: character.sql

package sqlc

import (
	"context"
	"log"
)

const FilteredCharactersByKeyset = `-- name: FilteredCharactersByKeyset :many
SELECT id, carac_simpl, carac_trad, pinyin, meaning, category, carac_antonym, carac_similar, radical_list
FROM character
WHERE carac_simpl ILIKE '%' || $1::text || '%' AND (id > $2::int OR $2::int IS NULL)
ORDER BY id
LIMIT $3::int
`

type FilteredCharactersByKeysetParams struct {
	Carac    string `json:"carac"`
	ID       int32  `json:"id"`
	Pagesize int32  `json:"pagesize"`
}

func (q *Queries) FilteredCharactersByKeyset(ctx context.Context, arg FilteredCharactersByKeysetParams) ([]Character, error) {
	rows, err := q.db.Query(ctx, FilteredCharactersByKeyset, arg.Carac, arg.ID, arg.Pagesize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Character
	for rows.Next() {
		var i Character
		if err := rows.Scan(
			&i.ID,
			&i.CaracSimpl,
			&i.CaracTrad,
			&i.Pinyin,
			&i.Meaning,
			&i.Category,
			&i.CaracAntonym,
			&i.CaracSimilar,
			&i.RadicalList,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetCharacterbyCategory = `-- name: GetCharacterbyCategory :many
SELECT id, carac_simpl, carac_trad, pinyin, meaning, category, carac_antonym, carac_similar, radical_list FROM character
WHERE category = $1::text
`

func (q *Queries) GetCharacterbyCategory(ctx context.Context, category string) ([]Character, error) {
	rows, err := q.db.Query(ctx, GetCharacterbyCategory, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Character
	for rows.Next() {
		var i Character
		if err := rows.Scan(
			&i.ID,
			&i.CaracSimpl,
			&i.CaracTrad,
			&i.Pinyin,
			&i.Meaning,
			&i.Category,
			&i.CaracAntonym,
			&i.CaracSimilar,
			&i.RadicalList,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetCharacterbyID = `-- name: GetCharacterbyID :one
SELECT id, carac_simpl, carac_trad, pinyin, meaning, category, carac_antonym, carac_similar, radical_list FROM character
WHERE id = $1::int32 LIMIT 1
`

func (q *Queries) GetCharacterbyID(ctx context.Context, id int) (Character, error) {
	row := q.db.QueryRow(ctx, GetCharacterbyID, id)
	var i Character
	err := row.Scan(
		&i.ID,
		&i.CaracSimpl,
		&i.CaracTrad,
		&i.Pinyin,
		&i.Meaning,
		&i.Category,
		&i.CaracAntonym,
		&i.CaracSimilar,
		&i.RadicalList,
	)
	return i, err
}

const ListCharactersByKeyset = `-- name: ListCharactersByKeyset :many
SELECT id, carac_simpl, carac_trad, pinyin, meaning, category, carac_antonym, carac_similar, radical_list
FROM character
WHERE (id > $1::int OR $1::int IS NULL)
ORDER BY id
LIMIT $2::int
`

type ListCharactersByKeysetParams struct {
	ID       int32 `json:"id"`
	Pagesize int32 `json:"pagesize"`
}

func (q *Queries) ListCharactersByKeyset(ctx context.Context, arg ListCharactersByKeysetParams) ([]Character, error) {
	rows, err := q.db.Query(ctx, ListCharactersByKeyset, arg.ID, arg.Pagesize)

	log.Printf("query passée")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Character
	log.Printf("item initialisé")

	for rows.Next() {
		var i Character
		if err := rows.Scan(
			&i.ID,
			&i.CaracSimpl,
			&i.CaracTrad,
			&i.Pinyin,
			&i.Meaning,
			&i.Category,
			&i.CaracAntonym,
			&i.CaracSimilar,
			&i.RadicalList,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
		log.Printf("je vais voir les tiems")
		log.Print(items)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	log.Printf("ListCharacterbyKeyset: %v items retrived", len(items))
	return items, nil
}
