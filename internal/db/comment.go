package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/QQuinn03/go_api/internal/comment"
	uuid "github.com/satori/go.uuid"
)

type Commentrow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func Convertcomment(cr Commentrow) comment.Comment {
	return comment.Comment{
		ID:     cr.ID,
		Slug:   cr.Slug.String,
		Body:   cr.Body.String,
		Author: cr.Author.String,
	}
}
func (d *Database) GetComment(ctx context.Context, id string) (comment.Comment, error) {
	var cmt Commentrow
	//row:=d.Client.QueryRowContext
	row := d.Client.QueryRowContext(ctx,
		`SELECT id,slug,body,author FROM comment WHERE id=$1`, id)
	// the pq driver for Postgres requires a placeholder like $1 instead of ?.
	err := row.Scan(&cmt.ID, &cmt.Slug, &cmt.Body, &cmt.Author)
	if err != nil {
		fmt.Println("error retivering data")
		return comment.Comment{}, err
	}

	Cmt := Convertcomment(cmt)
	return Cmt, nil

}

// Named queries can also use structs.  Their bind names follow the same rules
// as the name -> db mapping, so struct fields are lowercased and the `db` tag
// is taken into consideration.
func (d *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.NewV4().String()
	cmtRow := Commentrow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	queryResult, err := d.Client.NamedQueryContext(ctx, `INSERT INTO comment(id,slug,body,author) VALUES(:id,:slug,:body,:author)`, cmtRow)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}
	if err := queryResult.Close(); err != nil {
		return comment.Comment{}, err
	}

	return cmt, nil
}

func (d *Database) DeleteComment(ctx context.Context, id string) error {
	// _, err := d.Client.NamedQueryContext(ctx, `DELETE FROM comment WHERE ID=$1`, id)
	// if err != nil {
	// 	return err
	// }
	// return nil
	_, err := d.Client.NamedQueryContext(ctx, `DELETE FROM comment WHERE id=$1`, id)
	if err != nil {
		return fmt.Errorf("failto delete the comment:%v", err)

	}
	return nil
}

func (d *Database) UpdateComment(ctx context.Context, id string, cmt comment.Comment) (comment.Comment, error) {
	cmtRow := Commentrow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}
	updateRow, err := d.Client.NamedQueryContext(ctx,
		`UPDATE comment
	    SET id=:id,
	    slug=:slug,
		body=:body,
		author=:author`,
		cmtRow)
	if err != nil {
		return comment.Comment{}, err
	}
	if err := updateRow.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close table: %w", err)
	}
	return Convertcomment(cmtRow), nil

}
