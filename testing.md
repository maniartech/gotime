# ManiarTech® Testing Guidelines for Golang Projects ⧖

> **Professional testing standards for Go projects** - Ensuring consistent, high-quality tests that are readable, reliable, and effective.

These are the official ManiarTech® testing guidelines for testing any project written in Golang. All contributors must follow these best practices to ensure consistent, high-quality tests that are easy to read, reliable, and effective across the entire codebase.

---

## ◉ File & Package Structure

- Test files **must end with** `_test.go`.
- Test files must be placed in the **same directory** as the file being tested.
- You may use **either** of the following package strategies:

  ### White-box Testing (Access to Internal Code)
  - Use the **same package** name as the code under test.
  - Example:
    ```go
    // code.go
    package user
    // code_test.go
    package user
    ```
  - **✓ Use when** you want to test **unexported functions, types, or constants**.

  ### Black-box Testing (Public API Only)
  - Use the **`_test` suffix** in the package name.
  - Example:
    ```go
    package user_test
    ```
  - **✓ Use when** you want to test the **public API only**, simulating usage by external code.
  - **⟲ Preferred for most cases** to ensure loose coupling and clean design.

- If you're testing multiple packages, **replace `user`** with the appropriate package name from the current module.

- For **integration tests**, consider creating a `/test/` directory at the root level of the repo:
