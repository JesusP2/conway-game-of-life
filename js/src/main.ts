let prevTickTime = 0;
const squareSize = 20;
const width = 800;
const height = 800;
const columns = width / squareSize;
const rows = height / squareSize;
const FRAMES = 5;
let state = Array.from({ length: columns }, () =>
  Array.from({ length: rows }, () => {
    if (Math.random() > 0.8) {
      return true;
    }
    return false;
  }),
);
function setup() {
  const board = document.querySelector<HTMLCanvasElement>("#board");
  const ctx = board?.getContext("2d");
  if (!board || !ctx) return;
  board.width = width;
  board.height = height;
  ctx.fillStyle = "#e5e7eb";
  ctx.scale(squareSize, squareSize);
  requestAnimationFrame(tick);
}

function drawBoard(ctx: CanvasRenderingContext2D) {
  ctx.fillStyle = "#e5e7eb";
  ctx.scale(1 / squareSize, 1 / squareSize);
  for (let x = 1; x < columns; x++) {
    ctx.fillRect(x * squareSize, 0, 1, columns * squareSize);
  }
  for (let y = 1; y < rows; y++) {
    ctx.fillRect(0, y * squareSize, rows * squareSize, 1);
  }
  ctx.scale(squareSize, squareSize);
}

function tick(ms: number) {
  if (ms - prevTickTime < 1000 / FRAMES) {
    requestAnimationFrame(tick);
    return;
  }
  prevTickTime = ms;
  const board = document.querySelector<HTMLCanvasElement>("#board");
  const ctx = board?.getContext("2d");
  if (!board || !ctx) return;
  for (let y = 0; y < columns; y++) {
    for (let x = 0; x < rows; x++) {
      if (state[y][x]) {
        ctx.fillStyle = "#e5e7eb";
        ctx.fillRect(x, y, 1, 1);
      } else {
        ctx.fillStyle = "black";
        ctx.fillRect(x, y, 1, 1);
      }
    }
  }
  drawBoard(ctx);
  state = nextGeneration();
  requestAnimationFrame(tick);
}
function nextGeneration() {
  const nextState = state.map((row) => [...row]);
  for (let y = 0; y < columns; y++) {
    for (let x = 0; x < rows; x++) {
      const alive = getNeighbours(x, y).filter((neighbour) => neighbour);
      if (alive.length < 2 || alive.length > 3) {
        nextState[y][x] = false;
      } else if (alive.length === 3) {
        nextState[y][x] = true;
      }
    }
  }
  return nextState;
}

function getNeighbours(x: number, y: number) {
  const neighbours = [];
  for (let i = -1; i <= 1; i++) {
    for (let j = -1; j <= 1; j++) {
      if (!(i === 0 && j === 0)) {
        const value = state.at(y + i)?.at(x + j);
        if (typeof value === "boolean") {
          neighbours.push(value);
        }
      }
    }
  }
  return neighbours;
}

setup();
