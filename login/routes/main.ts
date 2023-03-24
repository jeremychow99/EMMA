import express from "express";
import 'express-async-errors';
import { register, login } from '../controllers/auth'
import getUsers from "../controllers/getUsers";

const router = express.Router()

router.get('/all', getUsers);
router.post('/register', register);
router.post('/login', login)
export default router;