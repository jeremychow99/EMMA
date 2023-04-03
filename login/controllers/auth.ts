import { User } from '../models/user';
import StatusCodes from 'http-status-codes';
import { Request, Response } from 'express';
import UnauthenticatedError from '../errors/unauthenticated';
import ResourceAlreadyExistsError from '../errors/alreadyExists';

const register = async (req: Request, res: Response) => {
  const { name, password, email, phone, role } = req.body
  if (name.length < 6 || password.length < 6) {
    throw new Error("Both Name and Password must be at least 6 characters long")
  }

  const checkUser = await User.findOne({ name })
  if (checkUser) {
    throw new ResourceAlreadyExistsError('Username already exists')
  }


  const user = await User.create({ ...req.body })
  const token = user.createJWT()
  res.status(StatusCodes.CREATED).json({ user: { userId: user._id, name: user.name, role: user.role, phonenumber: user.phone, email: user.email }, token })


}

const login = async (req: Request, res: Response) => {

  const { name, password } = req.body

  if (!name || !password) {
    throw new Error('Please provide email and password')
  }

  const user = await User.findOne({ name })
  if (!user) {
    throw new UnauthenticatedError('Invalid Credentials')
  }

  const isPasswordCorrect = await user.comparePassword(password)
  if (!isPasswordCorrect) {
    throw new UnauthenticatedError('Invalid Credentials')
  }
  const token = user.createJWT()
  res.status(StatusCodes.OK).json({ user: { userId: user._id, name: user.name, role: user.role, phonenumber: user.phone, email: user.email }, token })
}

export {
  register,
  login
}