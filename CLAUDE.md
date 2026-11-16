# Reading List Logic Port

Port the logic of the Reading List from TypeScript to Python 3. Read `README.md` and `SPEC.md` before you do anything.

I work **only** in the folder `python`. The folder `typescript` is the reference and read-only. The file `fixtures/cases.json` is shared by all languages and read-only.

## Commands

- Python, in `python`: `python3 -m unittest` (on Windows: `python -m unittest`)
- Reference, in `typescript`: `node --test`

## Rules

- **Do not change** the tests and the fixtures: `fixtures/cases.json`, `python/test_books.py`, `java/src/BooksTest.java`, `java/src/Json.java`, `go/books_test.go`, `rust/tests/cases.rs`, `typescript/*`. If a test fails, change the code. If you think a test or the fixtures are wrong, stop and ask me.
- Change only the source file of my language: `python/books.py`, `java/src/Books.java` (and `Book.java` only if needed), `go/books.go`, `rust/src/lib.rs`. Keep the names and signatures of the skeleton, the tests call them.
- **No new dependencies.** Standard library only. No library for lowercasing, trimming, searching or JSON. Do not change `Cargo.toml` or `go.mod`, do not add a `pip install`, a Maven or Gradle file or a build tool.
- Write the port in my language only. No other language, no script that generates code.
- **Plan mode first.** Explain the plan in terms of my language (the types, the standard library functions, what is different from TypeScript) and wait until I approve it. Then work in small steps: `STATUS_LABELS`, `filterBooks`, `validateBook`, with a test run after each.
- Trim the query **first**, then lower-case it, and only then use it (see `SPEC.md`).
- Keep the code simple and readable. A plain loop is fine. No clever tricks I can not review, and no `unsafe`.
- When you use a construct of the language that a TypeScript developer does not know, explain it in one sentence in your answer and compare it with TypeScript.
- Run the tests before you say a task is done, and show me the output. If you could not run them, say so.
- Do not commit unless I ask you to. I commit after I read the diff.

## Language rules

### Python 3

- Python 3.9 compatible. Type hints on all functions, `dataclass` for the book, standard library only.
