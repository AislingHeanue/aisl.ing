import React from "react";

function Lines(r, angle, n, rootx, rooty, rotation) {
  const lines = makeCoordinates(n, angle, r, rootx, rooty, rotation);

  return { lines: lines };
}

function makeCoordinates(n, angle, r, rootx, rooty, rotation) {
  const firstElement = [rootx, rooty].join(",");
  const firstAngle = rotation;
  let lines = [];
  for (let i = 2; i < n + 1; i++) {
    let currentAngle = firstAngle;
    let currentx = rootx,
      currenty = rooty;
    let newPolyline = [firstElement];

    let sequence = CollatzSequence(i);
    let lastNo = -1;
    for (let element of sequence) {
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

function CollatzSequence(n) {
  let sequence = [n];
  while (n != 1) {
    n = n % 2 == 0 ? n / 2 : (3 * n + 1) / 2;
    sequence.push(n);
  }
  return sequence.reverse();
}

function getCirclesFromLines(lists) {
  const list = lists;
  var dict = {};
  var res = [];
  for (var i = 0; i < list.length; i++) {
    for (let item of list[i]) {
      // console.log(item)
      if (!(item in dict)) {
        dict[item] = true;
        res.push(item);
      }
    }
  }
  return res;
}

export default Lines;
