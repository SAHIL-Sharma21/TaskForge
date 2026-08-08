import {initContract} from '@ts-rest/core';
import {healthContract} from './health.js';
import {todoContract} from "./todo.js";
import {categoryContact} from "./category.js";
import {commentContract} from "./comment.js";

const c = initContract();

export const apiContract = c.router({
    Health: healthContract,
    Todo: todoContract,
    Category: categoryContact,
    Comment: commentContract,
});
