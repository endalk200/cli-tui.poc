package internal

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
)

type GitCLI struct {
	repo *git.Repository
	path string
}

type ErrNotAGitRepository struct {
	Path string
}

func (e ErrNotAGitRepository) Error() string {
	return fmt.Sprintf("git: %s is not a git repository", e.Path)
}

type ErrUnkwownGitIssue struct {
	Message string
}

func (e ErrUnkwownGitIssue) Error() string {
	return fmt.Sprintf("git: unknown git issue: %s", e.Message)
}

func NewGitClient(repoPath string) (*GitCLI, error) {
	repo, err := git.PlainOpen(repoPath)

	if err != nil {
		if errors.Is(err, git.ErrRepositoryNotExists) {
			return nil, ErrNotAGitRepository{
				Path: repoPath,
			}
		}

		return nil, ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}

	return &GitCLI{repo: repo, path: repoPath}, nil
}

func (g *GitCLI) StagedFiles() ([]string, error) {
	workTree, err := g.repo.Worktree()
	if err != nil {
		return nil, ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}

	status, err := workTree.Status()
	if err != nil {
		return nil, ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}

	var stagedFiles []string
	for path, s := range status {
		if s.Staging != git.Unmodified { // something staged
			stagedFiles = append(stagedFiles, path)
		}
	}

	return stagedFiles, nil
}

func (g *GitCLI) ModifiedFiles() ([]string, error) {
	workTree, err := g.repo.Worktree()
	if err != nil {
		return nil, ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}

	status, err := workTree.Status()
	if err != nil {
		return nil, ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}

	var modifiedFiles []string
	for path, s := range status {
		if s.Worktree != git.Unmodified {
			modifiedFiles = append(modifiedFiles, path)
		}
	}
	return modifiedFiles, nil
}

func (g *GitCLI) AddedFiles() ([]string, error) {
	workTree, err := g.repo.Worktree()
	if err != nil {
		return nil, ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}

	status, err := workTree.Status()
	if err != nil {
		return nil, ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}

	var addedFiles []string
	for path, s := range status {
		if s.Worktree == git.Added {
			addedFiles = append(addedFiles, path)
		}
	}
	return addedFiles, nil
}

func (g *GitCLI) DeletedFiles() ([]string, error) {
	workTree, err := g.repo.Worktree()
	if err != nil {
		return nil, ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}

	status, err := workTree.Status()
	if err != nil {
		return nil, ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}

	var deletedFiles []string
	for path, s := range status {
		if s.Worktree == git.Deleted {
			deletedFiles = append(deletedFiles, path)
		}
	}

	return deletedFiles, nil
}

func (g *GitCLI) RenamedFiles() ([]string, error) {
	workTree, err := g.repo.Worktree()
	if err != nil {
		return nil, ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}

	status, err := workTree.Status()
	if err != nil {
		return nil, ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}

	var renamedFiles []string
	for path, s := range status {
		if s.Worktree == git.Renamed {
			renamedFiles = append(renamedFiles, path)
		}
	}

	return renamedFiles, nil
}

func (g *GitCLI) UntrackedFiles() ([]string, error) {
	workTree, err := g.repo.Worktree()
	if err != nil {
		return nil, ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}

	status, err := workTree.Status()
	if err != nil {
		return nil, ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}

	var untrackedFiles []string
	for path, s := range status {
		if s.Worktree == git.Untracked {
			untrackedFiles = append(untrackedFiles, path)
		}
	}

	return untrackedFiles, nil
}

func (g *GitCLI) GetStagedFilesDiff(stagedFiles []string) (string, error) {
	var diff string
	for _, file := range stagedFiles {
		cmd := exec.Command("git", "diff", file)
		cmd.Dir = g.path
		out, err := cmd.CombinedOutput()
		if err != nil {
			return "", ErrUnkwownGitIssue{Message: err.Error()}
		}

		diff += string(out)
	}

	return string(diff), nil
}

func (g *GitCLI) CurrentBranch() (string, error) {
	headRef, err := g.repo.Head()
	if err != nil {
		return "", ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}
	name := headRef.Name().String()
	if strings.HasPrefix(name, "refs/heads/") {
		return strings.TrimPrefix(name, "refs/heads/"), nil
	}
	return headRef.Hash().String()[:12], nil
}

func (g *GitCLI) Commit(message string) error {
	workTree, err := g.repo.Worktree()
	if err != nil {
		return ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}

	author := &object.Signature{
		Name:  "endalk200",
		Email: "eb808826@gmail.com",
		When:  time.Now(),
	}

	commitHash, err := workTree.Commit(message, &git.CommitOptions{
		Author:    author,
		Committer: author,
		All:       false,
	})
	if err != nil {
		return ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}

	commitObj, err := g.repo.CommitObject(commitHash)
	if err != nil {
		return ErrUnkwownGitIssue{
			Message: err.Error(),
		}
	}

	fmt.Println("✅ Commit created successfully!")
	fmt.Printf("  📝 Hash: %s\n", commitObj.Hash.String()[:7])
	fmt.Printf("  👤 Author: %s <%s>\n", commitObj.Author.Name, commitObj.Author.Email)
	if commitObj.Committer != commitObj.Author {
		fmt.Printf("  ✉️  Committer: %s <%s>\n", commitObj.Committer.Name, commitObj.Committer.Email)
	}
	fmt.Printf("  🕐 Date: %s\n", commitObj.Author.When.Format(time.RFC1123))
	fmt.Printf("  📄 Message: %s\n", message)
	return nil
}
