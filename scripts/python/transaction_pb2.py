# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x11transaction.proto\x12\x06pactus\"c\n\x15GetTransactionRequest\x12\x0e\n\x02id\x18\x01 \x01(\x0cR\x02id\x12:\n\tverbosity\x18\x02 \x01(\x0e\x32\x1c.pactus.TransactionVerbosityR\tverbosity\"\x95\x01\n\x16GetTransactionResponse\x12!\n\x0c\x62lock_height\x18\x0c \x01(\rR\x0b\x62lockHeight\x12\x1d\n\nblock_time\x18\r \x01(\rR\tblockTime\x12\x39\n\x0btransaction\x18\x03 \x01(\x0b\x32\x17.pactus.TransactionInfoR\x0btransaction\"d\n\x13\x43\x61lculateFeeRequest\x12\x16\n\x06\x61mount\x18\x01 \x01(\x03R\x06\x61mount\x12\x35\n\x0bpayloadType\x18\x02 \x01(\x0e\x32\x13.pactus.PayloadTypeR\x0bpayloadType\"(\n\x14\x43\x61lculateFeeResponse\x12\x10\n\x03\x66\x65\x65\x18\x01 \x01(\x03R\x03\x66\x65\x65\":\n\x1b\x42roadcastTransactionRequest\x12\x1b\n\tsigned_tx\x18\x01 \x01(\x0cR\x08signedTx\".\n\x1c\x42roadcastTransactionResponse\x12\x0e\n\x02id\x18\x02 \x01(\x0cR\x02id\"\xb1\x01\n GetRawTransferTransactionRequest\x12\x1b\n\tlock_time\x18\x01 \x01(\rR\x08lockTime\x12\x16\n\x06sender\x18\x02 \x01(\tR\x06sender\x12\x1a\n\x08receiver\x18\x03 \x01(\tR\x08receiver\x12\x16\n\x06\x61mount\x18\x04 \x01(\x03R\x06\x61mount\x12\x10\n\x03\x66\x65\x65\x18\x05 \x01(\x03R\x03\x66\x65\x65\x12\x12\n\x04memo\x18\x06 \x01(\tR\x04memo\"\xca\x01\n\x1cGetRawBondTransactionRequest\x12\x1b\n\tlock_time\x18\x01 \x01(\rR\x08lockTime\x12\x16\n\x06sender\x18\x02 \x01(\tR\x06sender\x12\x1a\n\x08receiver\x18\x03 \x01(\tR\x08receiver\x12\x14\n\x05stake\x18\x04 \x01(\x03R\x05stake\x12\x1d\n\npublic_key\x18\x05 \x01(\tR\tpublicKey\x12\x10\n\x03\x66\x65\x65\x18\x06 \x01(\x03R\x03\x66\x65\x65\x12\x12\n\x04memo\x18\x07 \x01(\tR\x04memo\"~\n\x1eGetRawUnBondTransactionRequest\x12\x1b\n\tlock_time\x18\x01 \x01(\rR\x08lockTime\x12+\n\x11validator_address\x18\x03 \x01(\tR\x10validatorAddress\x12\x12\n\x04memo\x18\x04 \x01(\tR\x04memo\"\xd3\x01\n GetRawWithdrawTransactionRequest\x12\x1b\n\tlock_time\x18\x01 \x01(\rR\x08lockTime\x12+\n\x11validator_address\x18\x03 \x01(\tR\x10validatorAddress\x12\'\n\x0f\x61\x63\x63ount_address\x18\x04 \x01(\tR\x0e\x61\x63\x63ountAddress\x12\x10\n\x03\x66\x65\x65\x18\x05 \x01(\x03R\x03\x66\x65\x65\x12\x16\n\x06\x61mount\x18\x06 \x01(\x03R\x06\x61mount\x12\x12\n\x04memo\x18\x07 \x01(\tR\x04memo\"D\n\x19GetRawTransactionResponse\x12\'\n\x0fraw_transaction\x18\x01 \x01(\x0cR\x0erawTransaction\"]\n\x0fPayloadTransfer\x12\x16\n\x06sender\x18\x01 \x01(\tR\x06sender\x12\x1a\n\x08receiver\x18\x02 \x01(\tR\x08receiver\x12\x16\n\x06\x61mount\x18\x03 \x01(\x03R\x06\x61mount\"W\n\x0bPayloadBond\x12\x16\n\x06sender\x18\x01 \x01(\tR\x06sender\x12\x1a\n\x08receiver\x18\x02 \x01(\tR\x08receiver\x12\x14\n\x05stake\x18\x03 \x01(\x03R\x05stake\"B\n\x10PayloadSortition\x12\x18\n\x07\x61\x64\x64ress\x18\x01 \x01(\tR\x07\x61\x64\x64ress\x12\x14\n\x05proof\x18\x02 \x01(\x0cR\x05proof\"-\n\rPayloadUnbond\x12\x1c\n\tvalidator\x18\x01 \x01(\tR\tvalidator\"M\n\x0fPayloadWithdraw\x12\x12\n\x04\x66rom\x18\x01 \x01(\tR\x04\x66rom\x12\x0e\n\x02to\x18\x02 \x01(\tR\x02to\x12\x16\n\x06\x61mount\x18\x03 \x01(\x03R\x06\x61mount\"\xab\x04\n\x0fTransactionInfo\x12\x0e\n\x02id\x18\x01 \x01(\x0cR\x02id\x12\x12\n\x04\x64\x61ta\x18\x02 \x01(\x0cR\x04\x64\x61ta\x12\x18\n\x07version\x18\x03 \x01(\x05R\x07version\x12\x1b\n\tlock_time\x18\x04 \x01(\rR\x08lockTime\x12\x14\n\x05value\x18\x05 \x01(\x03R\x05value\x12\x10\n\x03\x66\x65\x65\x18\x06 \x01(\x03R\x03\x66\x65\x65\x12\x35\n\x0bpayloadType\x18\x07 \x01(\x0e\x32\x13.pactus.PayloadTypeR\x0bpayloadType\x12\x35\n\x08transfer\x18\x1e \x01(\x0b\x32\x17.pactus.PayloadTransferH\x00R\x08transfer\x12)\n\x04\x62ond\x18\x1f \x01(\x0b\x32\x13.pactus.PayloadBondH\x00R\x04\x62ond\x12\x38\n\tsortition\x18  \x01(\x0b\x32\x18.pactus.PayloadSortitionH\x00R\tsortition\x12/\n\x06unbond\x18! \x01(\x0b\x32\x15.pactus.PayloadUnbondH\x00R\x06unbond\x12\x35\n\x08withdraw\x18\" \x01(\x0b\x32\x17.pactus.PayloadWithdrawH\x00R\x08withdraw\x12\x12\n\x04memo\x18\x08 \x01(\tR\x04memo\x12\x1d\n\npublic_key\x18\t \x01(\tR\tpublicKey\x12\x1c\n\tsignature\x18\n \x01(\x0cR\tsignatureB\t\n\x07payload*\x83\x01\n\x0bPayloadType\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x14\n\x10TRANSFER_PAYLOAD\x10\x01\x12\x10\n\x0c\x42OND_PAYLOAD\x10\x02\x12\x15\n\x11SORTITION_PAYLOAD\x10\x03\x12\x12\n\x0eUNBOND_PAYLOAD\x10\x04\x12\x14\n\x10WITHDRAW_PAYLOAD\x10\x05*B\n\x14TransactionVerbosity\x12\x14\n\x10TRANSACTION_DATA\x10\x00\x12\x14\n\x10TRANSACTION_INFO\x10\x01\x32\xa8\x05\n\x0bTransaction\x12O\n\x0eGetTransaction\x12\x1d.pactus.GetTransactionRequest\x1a\x1e.pactus.GetTransactionResponse\x12I\n\x0c\x43\x61lculateFee\x12\x1b.pactus.CalculateFeeRequest\x1a\x1c.pactus.CalculateFeeResponse\x12\x61\n\x14\x42roadcastTransaction\x12#.pactus.BroadcastTransactionRequest\x1a$.pactus.BroadcastTransactionResponse\x12h\n\x19GetRawTransferTransaction\x12(.pactus.GetRawTransferTransactionRequest\x1a!.pactus.GetRawTransactionResponse\x12`\n\x15GetRawBondTransaction\x12$.pactus.GetRawBondTransactionRequest\x1a!.pactus.GetRawTransactionResponse\x12\x64\n\x17GetRawUnBondTransaction\x12&.pactus.GetRawUnBondTransactionRequest\x1a!.pactus.GetRawTransactionResponse\x12h\n\x19GetRawWithdrawTransaction\x12(.pactus.GetRawWithdrawTransactionRequest\x1a!.pactus.GetRawTransactionResponseBF\n\x12pactus.transactionZ0github.com/pactus-project/pactus/www/grpc/pactusb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'transaction_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\022pactus.transactionZ0github.com/pactus-project/pactus/www/grpc/pactus'
  _PAYLOADTYPE._serialized_start=2268
  _PAYLOADTYPE._serialized_end=2399
  _TRANSACTIONVERBOSITY._serialized_start=2401
  _TRANSACTIONVERBOSITY._serialized_end=2467
  _GETTRANSACTIONREQUEST._serialized_start=29
  _GETTRANSACTIONREQUEST._serialized_end=128
  _GETTRANSACTIONRESPONSE._serialized_start=131
  _GETTRANSACTIONRESPONSE._serialized_end=280
  _CALCULATEFEEREQUEST._serialized_start=282
  _CALCULATEFEEREQUEST._serialized_end=382
  _CALCULATEFEERESPONSE._serialized_start=384
  _CALCULATEFEERESPONSE._serialized_end=424
  _BROADCASTTRANSACTIONREQUEST._serialized_start=426
  _BROADCASTTRANSACTIONREQUEST._serialized_end=484
  _BROADCASTTRANSACTIONRESPONSE._serialized_start=486
  _BROADCASTTRANSACTIONRESPONSE._serialized_end=532
  _GETRAWTRANSFERTRANSACTIONREQUEST._serialized_start=535
  _GETRAWTRANSFERTRANSACTIONREQUEST._serialized_end=712
  _GETRAWBONDTRANSACTIONREQUEST._serialized_start=715
  _GETRAWBONDTRANSACTIONREQUEST._serialized_end=917
  _GETRAWUNBONDTRANSACTIONREQUEST._serialized_start=919
  _GETRAWUNBONDTRANSACTIONREQUEST._serialized_end=1045
  _GETRAWWITHDRAWTRANSACTIONREQUEST._serialized_start=1048
  _GETRAWWITHDRAWTRANSACTIONREQUEST._serialized_end=1259
  _GETRAWTRANSACTIONRESPONSE._serialized_start=1261
  _GETRAWTRANSACTIONRESPONSE._serialized_end=1329
  _PAYLOADTRANSFER._serialized_start=1331
  _PAYLOADTRANSFER._serialized_end=1424
  _PAYLOADBOND._serialized_start=1426
  _PAYLOADBOND._serialized_end=1513
  _PAYLOADSORTITION._serialized_start=1515
  _PAYLOADSORTITION._serialized_end=1581
  _PAYLOADUNBOND._serialized_start=1583
  _PAYLOADUNBOND._serialized_end=1628
  _PAYLOADWITHDRAW._serialized_start=1630
  _PAYLOADWITHDRAW._serialized_end=1707
  _TRANSACTIONINFO._serialized_start=1710
  _TRANSACTIONINFO._serialized_end=2265
  _TRANSACTION._serialized_start=2470
  _TRANSACTION._serialized_end=3150
# @@protoc_insertion_point(module_scope)