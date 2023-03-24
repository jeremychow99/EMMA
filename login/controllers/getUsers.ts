import { User } from '../models/user';
import StatusCodes from 'http-status-codes';
import { Request, Response } from 'express';

const getUsers = async (req: Request, res: Response) => {
    const users = await User.find({});
    res.status(StatusCodes.OK).json({ users: users })
}

export default getUsers