from e2b import Sandbox
from e2b.sandbox.process.process_handle import PtySize


def test_pty(sandbox: Sandbox):
    def append_data(data:list, x: bytes):
        data.append(x.decode("utf-8"))

    terminal = sandbox.pty.create(
        PtySize(80, 24), envs={"ABC": "123"}, cwd="/"
    )

    sandbox.pty.send_stdin(terminal.pid, b"echo $ABC\n")
    sandbox.pty.send_stdin(terminal.pid, b"pwd\n")
    sandbox.pty.send_stdin(terminal.pid, b"logout\n")

    output = []
    result = terminal.wait(on_pty=lambda x: append_data(output, x))
    assert result.exit_code == 0

    assert "123" in '\n'.join(output)