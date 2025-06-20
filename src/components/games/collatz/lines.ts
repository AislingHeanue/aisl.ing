import { State } from "./store";

export default function Lines({ r, angle, n, rootX, rootY, rotation }: State) {

  const firstElement = [rootX, rootY].join(",");
  const firstAngle = rotation;
  const lines = [];
  for (let i = 2; i < n + 1; i++) {
    let currentAngle = firstAngle;
    let currentx = rootX,
      currenty = rootY;
    const newPolyline = [firstElement];

    const sequence = CollatzSequence(i);
    let lastNo = -1;
    for (const element of sequence) {
      if (lastNo == -1) {
        lastNo = element;
        continue;
      }

      if (element == 2 * lastNo) {
        currentAngle -= angle;
      } else {
        currentAngle += angle;
      }
      lastNo = element;
      currentx += r * Math.cos(currentAngle);
      currenty -= r * Math.sin(currentAngle);
      // console.log(currentx)
      newPolyline.push([currentx, currenty].join(","));
    }
    lines.push(newPolyline);
  }
  return lines;
}

function CollatzSequence(n: number) {
  const sequence = [n];
  while (n != 1) {
    n = n % 2 == 0 ? n / 2 : (3 * n + 1) / 2;
    sequence.push(n);
  }
  return sequence.reverse();
}

// function getCirclesFromLines(lists: ) {
//   const list = lists;
//   var dict = {};
//   var res = [];
//   for (var i = 0; i < list.length; i++) {
//     for (let item of list[i]) {
//       // console.log(item)
//       if (!(item in dict)) {
//         dict[item] = true;
//         res.push(item);
//       }
//     }
//   }
//   return res;
// }

