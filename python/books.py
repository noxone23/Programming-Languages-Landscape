"""The Reading List logic, ported from TypeScript (see ../typescript/books.ts and ../SPEC.md)."""

from dataclasses import dataclass
from typing import Dict, List, Optional


@dataclass
class Book:
    id: str
    title: str
    author: str
    status: str  # "want-to-read", "reading" or "done"
    rating: Optional[int]  # 1 to 5, None means "not rated"
    notes: str


STATUS_LABELS: Dict[str, str] = {
    "want-to-read": "Want to read",
    "reading": "Reading",
    "done": "Done",
}


def filter_books(books: List[Book], status: str, query: str) -> List[Book]:
    """status is "want-to-read", "reading", "done" or "all"."""
    normalized_query = query.strip().lower()
    result: List[Book] = []
    for book in books:
        status_matches = status == "all" or book.status == status
        if status_matches and normalized_query in book.title.lower():
            result.append(book)
    return result


def validate_book(title: str, author: str) -> Dict[str, str]:
    """Returns an error message per invalid field. An empty dict means valid."""
    errors: Dict[str, str] = {}
    if title.strip() == "":
        errors["title"] = "Title is required"
    if author.strip() == "":
        errors["author"] = "Author is required"
    return errors
