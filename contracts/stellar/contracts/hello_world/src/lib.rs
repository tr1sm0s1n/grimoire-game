#![no_std]
use soroban_sdk::{contract, contractimpl, Env, Address, String};
use core::u64;


#[contract]
pub struct NFTContract;

#[contractimpl]
impl NFTContract {

    pub fn mint(e: &Env, nft_id: String, addr: Address, score: u64) {
        let key = nft_id;
        let value = (addr, score);
        e.storage().persistent().set(&key, &value);
    }

    pub fn get_nft_details(env: Env, nft_id: String) -> (Address, u64) {
        let key = nft_id;
        match env.storage().persistent().get::<String, (Address, u64)>(&key) {
            Some((address, score)) => (address, score),
            None => panic!("NFT ID not found"),
        }
    }

    pub fn transfer(env: Env, nft_id: String, to: Address) {
        let key = nft_id;
        match env.storage().persistent().get::<String, (Address, u64)>(&key) {
            Some((_, score)) => {
                let value = (to, score);
                env.storage().persistent().set(&key, &value);
            },
            None => panic!("NFT ID not found"),
        }
    }

    pub fn update_score(env: Env, nft_id: String, new_score: u64) {
        let key = nft_id;
        match env.storage().persistent().get::<String, (Address, u64)>(&key) {
            Some((address, _)) => {
                let value = (address, new_score);
                env.storage().persistent().set(&key, &value);
            },
            None => panic!("NFT ID not found"),
        }
    }
}