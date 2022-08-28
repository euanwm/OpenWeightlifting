import React, { useEffect, useState } from "react";
import * as ReactBootStrap from "react-bootstrap";
import axios from "axios";

const AxiosPost = () => {
  const [posts, setPosts] = useState({ blogs: [] });

  useEffect(() => {
    const fetchPostList = async () => {
      const { data } = await axios(
        "http://localhost:8000/api/leaderboard"
      );
      setPosts({ blogs: data });
      console.log(data);
    };
    fetchPostList().then();
  }, [setPosts]);

  return (
    <div>
      <ReactBootStrap.Table striped bordered hover>
        <thead>
          <tr>
            <th>Position</th>
            <th>Lifter</th>
            <th>Country</th>
            <th>Bodyweight</th>
            <th>Snatch 1</th>
            <th>Snatch 2</th>
            <th>Snatch 3</th>
            <th>Clean & Jerk 1</th>
            <th>Clean & Jerk 2</th>
            <th>Clean & Jerk 3</th>
            <th>Total</th>
          </tr>
        </thead>
        <tbody>
          {posts.blogs &&
            posts.blogs.map((item) => (
              <tr key={item.id}>
                <td>{item.id + 1}</td>
                <td>{item.lifter_name}</td>
                <td>{item.country}</td>
                <td>{item.bodyweight}</td>
                <td>{item.snatch_1}</td>
                <td>{item.snatch_2}</td>
                <td>{item.snatch_3}</td>
                <td>{item.cj_1}</td>
                <td>{item.cj_2}</td>
                <td>{item.cj_3}</td>
                <td>{item.total}</td>
              </tr>
            ))}
        </tbody>
      </ReactBootStrap.Table>
    </div>
  );
};

export default AxiosPost;
