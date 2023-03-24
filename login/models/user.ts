import mongoose from 'mongoose';
import bcrypt from 'bcrypt';
import jwt from 'jsonwebtoken'
import sanitizedConfig from '../utils/config'

const UserSchema: mongoose.Schema = new mongoose.Schema({
    name: {
        type: String,
        unique: true
    },
    password: { type: String },
    role: { type: String },
});

const iss = "test123";

UserSchema.pre('save', async function () {
    const salt = await bcrypt.genSalt(10);
    this.password = await bcrypt.hash(this.password, salt)
    if(this.role == ""){
        this.role == "TECHNICIAN";
    }

})

UserSchema.methods.createJWT = function () {
    return jwt.sign({ iss: iss }, sanitizedConfig.JWT_SECRET, { expiresIn: sanitizedConfig.JWT_LIFETIME })
}

UserSchema.methods.comparePassword = async function (candidatePassword: any) {
    const isMatch = await bcrypt.compare(candidatePassword, this.password)
    return isMatch
}

export const User = mongoose.model('User', UserSchema);